//TODO: correct license

package morello

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/syzkaller/pkg/config"
	"github.com/google/syzkaller/pkg/log"
	"github.com/google/syzkaller/pkg/osutil"
	"github.com/google/syzkaller/pkg/report"
	"github.com/google/syzkaller/sys/targets"
	"github.com/google/syzkaller/vm/vmimpl"
)

func init() {
	vmimpl.Register("morello", ctor, false)
}

type Config struct {
	MorelloBoxAddress  string `json:"mbox_address"`
	MorelloBoxUserName string `json:"mbox_username"`
	MorelloBoxSSHPort  int    `json:"mbox_port"`
	MorelloBoxSSHKey   string `json:"mbox_sshkey"`
	MorelloBoxWorkDir  string `json:"mbox_workdir"`
	RebootCommand      string `json:"reboot_command"`
}

type Pool struct {
	env *vmimpl.Env
	cfg *Config
}

type machine struct {
	sshuser        string
	address        string
	sshport        int
	sshkey         string
	forwardingPort int
}

type instance struct {
	cfg        *Config
	closed     chan bool
	debug      bool
	workdir    string
	os         string
	timeouts   targets.Timeouts
	merger     *vmimpl.OutputMerger
	writePipe  io.WriteCloser
	rpipe      io.ReadCloser
	morellobox machine
}

func ctor(env *vmimpl.Env) (vmimpl.Pool, error) {
	cfg := &Config{}
	if err := config.LoadData(env.Config, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse morello config: %v", err)
	}

	if cfg.MorelloBoxAddress == "" {
		return nil, fmt.Errorf("config param morelloBoxAddress is empty")
	}
	if cfg.RebootCommand == "" {
		return nil, fmt.Errorf("reboot command is empty")
	}
	if cfg.MorelloBoxWorkDir == "" {
		return nil, fmt.Errorf("config param morello box working directory is empty")
	}

	if cfg.MorelloBoxSSHPort == 0 {
		cfg.MorelloBoxSSHPort = 22
	}
	if cfg.MorelloBoxUserName == "" {
		cfg.MorelloBoxUserName = env.SSHUser
	}
	if cfg.MorelloBoxSSHKey == "" {
		cfg.MorelloBoxSSHKey = env.SSHKey
	}

	pool := &Pool{
		cfg: cfg,
		env: env,
	}
	return pool, nil
}

func (pool *Pool) Count() int {
	return 1 // no support for multiple Morello devices yet
}

func (pool *Pool) Create(workdir string, index int) (vmimpl.Instance, error) {
	inst := &instance{
		closed:   make(chan bool),
		cfg:      pool.cfg,
		debug:    pool.env.Debug,
		timeouts: pool.env.Timeouts,
		os:       pool.env.OS,
		workdir:  workdir,
	}
	inst.morellobox = machine{
		sshuser:        inst.cfg.MorelloBoxUserName,
		address:        inst.cfg.MorelloBoxAddress,
		sshport:        inst.cfg.MorelloBoxSSHPort,
		sshkey:         inst.cfg.MorelloBoxSSHKey,
		forwardingPort: 0,
	}
	var tee io.Writer
	if inst.debug {
		tee = os.Stdout
	}
	inst.merger = vmimpl.NewOutputMerger(tee)

	closeInst := inst
	defer func() {
		if closeInst != nil {
			closeInst.Close()
		}
	}()
	if err := inst.start_morellobox(); err != nil {
		return nil, err
	}

	closeInst = nil
	return inst, nil
}

func (inst *instance) start_morellobox() error {
	args := strings.Split(inst.cfg.RebootCommand, " ")
	if inst.debug {
		log.Logf(0, "running run command:%#v", args)
	}
	args = append(args, "")
	_, err := osutil.RunCmd(90*time.Second, "", args[0], args[1:]...)
	if err != nil {
		if inst.debug {
			log.Logf(0, "running reboot command failed:%v", err)
		}
	}
	if err := inst.waitForSSH(3*time.Minute, inst.morellobox); err == nil {
		return nil
	}
	return fmt.Errorf("Couldn't reboot")

}

func (inst *instance) isClosed() bool { //TODO: delete if not needed
	select {
	case <-inst.closed:
		return true
	default:
		return false
	}
}

func (inst *instance) runCommand(command string, target machine, timeout time.Duration) error {
	sshArgs := vmimpl.SSHArgs(inst.debug, target.sshkey, target.sshport)
	args := append([]string{"ssh"}, sshArgs...)
	args = append(args, target.fullAddress(), command)
	if inst.debug {
		log.Logf(0, "running run command:%#v", args)
	}
	_, err := osutil.RunCmd(timeout, "", args[0], args[1:]...)
	return err
}

func (target machine) fullAddress() string {
	return target.sshuser + "@" + target.address
}

func (inst *instance) waitForSSH(timeout time.Duration, target machine) error {
	if inst.debug {
		log.Logf(0, "Starting waiting for ssh, with target:%v", target.address)
	}
	err := vmimpl.WaitForSSH(false, timeout, inst.cfg.MorelloBoxAddress,
		target.sshkey, target.sshuser, inst.os, target.sshport, inst.merger.Err)
	return err
}

func (inst *instance) Forward(port int) (string, error) {
	if port == 0 {
		return "", fmt.Errorf("forward port is zero")
	}
	if inst.morellobox.forwardingPort != 0 {
		return "", fmt.Errorf("forward port is already set")
	}
	inst.morellobox.forwardingPort = port
	return fmt.Sprintf("localhost:%v", port), nil
}

func (inst *instance) startDiagnose() {

}

func (inst *instance) Diagnose(rep *report.Report) ([]byte, bool) {
	//TODO implement
	ret := []byte(fmt.Sprintf("No diagnose yet"))
	return ret, false
}

func (inst *instance) Close() {

	if inst.writePipe != nil {
		inst.writePipe.Close()
	}
	if inst.rpipe != nil {
		inst.rpipe.Close()
	}
	if inst.merger != nil {
		inst.merger.Wait()
	}
	close(inst.closed)
}

func (inst *instance) Copy(hostSrc string) (string, error) {
	basePath := filepath.Base(hostSrc)
	vmDst := filepath.Join(inst.cfg.MorelloBoxWorkDir, basePath)
	args := append(vmimpl.SCPArgs(inst.debug, inst.morellobox.sshkey, inst.morellobox.sshport), hostSrc,
		inst.morellobox.fullAddress()+":"+vmDst)
	if inst.debug {
		log.Logf(0, "running command: scp %#v", args)
	}
	_, err := osutil.RunCmd(10*time.Minute*inst.timeouts.Scale, "", "scp", args...)
	if err != nil {
		return "", err
	}
	return vmDst, nil
}

func (inst *instance) Run(timeout time.Duration, stop <-chan bool, command string) (
	<-chan []byte, <-chan error, error) {
	var err error
	var wpipe io.WriteCloser
	inst.rpipe, wpipe, err = osutil.LongPipe()
	if err != nil {
		return nil, nil, err
	}
	inst.merger.Add("ssh", inst.rpipe)

	//Needed while CTSRD-CHERI/cheribsd-ports/issues/9 is not fixed
	command = "env GODEBUG=asyncpreemptoff=1 " + command

	command = "sudo " + command
	command = "cd " + inst.cfg.MorelloBoxWorkDir + " ;" + command

	sshArgs := vmimpl.SSHArgsForward(inst.debug, inst.morellobox.sshkey, inst.morellobox.sshport, inst.morellobox.forwardingPort)
	args := append([]string{"ssh"}, sshArgs...)
	args = append(args, inst.morellobox.fullAddress(), command)

	if inst.debug {
		log.Logf(0, "running command:%#v", args)
	}
	cmd := osutil.Command(args[0], args[1:]...)
	cmd.Dir = inst.workdir
	cmd.Stdout = wpipe
	cmd.Stderr = wpipe
	if err := cmd.Start(); err != nil {
		wpipe.Close()
		return nil, nil, err
	}
	wpipe.Close()

	errc := make(chan error, 1)
	signal := func(err error) {
		select {
		case errc <- err:
		default:
		}
	}

	go func() {
		select {
		case <-time.After(timeout):
			signal(vmimpl.ErrTimeout)
		case <-stop:
			signal(vmimpl.ErrTimeout)
		case err := <-inst.merger.Err:
			cmd.Process.Kill()
			if cmdErr := cmd.Wait(); cmdErr == nil {
				// If the command exited successfully, we got EOF error from merger.
				// But in this case no error has happened and the EOF is expected.
				err = nil
			}
			signal(err)
			return
		}
		cmd.Process.Kill()
		cmd.Wait()
	}()
	return inst.merger.Output, errc, nil
}
