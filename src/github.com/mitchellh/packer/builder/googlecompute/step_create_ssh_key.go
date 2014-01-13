package googlecompute

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"code.google.com/p/go.crypto/ssh"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

// StepCreateSSHKey represents a Packer build step that generates SSH key pairs.
type StepCreateSSHKey int

// Run executes the Packer build step that generates SSH key pairs.
func (s *StepCreateSSHKey) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Creating temporary SSH key for instance...")
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		err := fmt.Errorf("Error creating temporary ssh key: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	priv_blk := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(priv),
	}

	pub, err := ssh.NewPublicKey(&priv.PublicKey)
	if err != nil {
		err := fmt.Errorf("Error creating temporary ssh key: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	state.Put("ssh_private_key", string(pem.EncodeToMemory(&priv_blk)))
	state.Put("ssh_public_key", string(ssh.MarshalAuthorizedKey(pub)))
	return multistep.ActionContinue
}

// Nothing to clean up. SSH keys are associated with a single GCE instance.
func (s *StepCreateSSHKey) Cleanup(state multistep.StateBag) {}
