package ngmanager

import (
	"encoding/json"
	"fmt"
)

func (a *Secret) UnmarshalJSON(data []byte) error {
	fmt.Println("inside unmarshal json")

	type Alias Secret

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	err = json.Unmarshal(aux.Spec, &a.SecretSpec)
	if err != nil {
		return err
	}

	switch a.SecretSpec.Type_ {
	case "SSHKeyPath":
		err = json.Unmarshal(aux.Spec, &a.SshKeyPathSpec)
	case "SSHKeyReference":
		err = json.Unmarshal(aux.Spec, &a.SshKeyReferenceSpec)
	case "SSHPassword":
		err = json.Unmarshal(aux.Spec, &a.SshPasswordSpec)
	case "SSHKerberosTGTKeyTabFile":
		err = json.Unmarshal(aux.Spec, &a.SshKerberosTgtKeyTabFileSpec)
	case "SSHKerberosTGTPassword":
		err = json.Unmarshal(aux.Spec, &a.SshKerberosTgtPasswordSpec)
	case "SecretFile":
		err = json.Unmarshal(aux.Spec, &a.SecretFileSpec)
	case "SecretText":
		err = json.Unmarshal(aux.Spec, &a.SecretTextSpec)
	case "WinRmTGTKeyTabFile":
		err = json.Unmarshal(aux.Spec, &a.WinRmTgtKeyTabFileSpec)
	case "WinRmTGTPassword":
		err = json.Unmarshal(aux.Spec, &a.WinRmTgtPasswordSpec)
	case "WinRmNTLM":
		err = json.Unmarshal(aux.Spec, &a.WinRmNtlmSpec)
	default:
		panic(fmt.Sprintf("unknown secret type %s", a.SecretSpec.Type_))
	}

	return err
}

func (a *Secret) MarshalJSON() ([]byte, error) {
	fmt.Println("inside marshal json")
	type Alias Secret

	var spec []byte
	var err error

	switch a.SecretSpec.Type_ {
	case "SSHKeyPath":
		spec, err = json.Marshal(a.SshKeyPathSpec)
	case "SSHKeyReference":
		spec, err = json.Marshal(a.SshKeyReferenceSpec)
	case "SSHPassword":
		spec, err = json.Marshal(a.SshPasswordSpec)
	case "SSHKerberosTGTKeyTabFile":
		spec, err = json.Marshal(a.SshKerberosTgtKeyTabFileSpec)
	case "SSHKerberosTGTPassword":
		spec, err = json.Marshal(a.SshKerberosTgtPasswordSpec)
	case "SecretFile":
		spec, err = json.Marshal(a.SecretFileSpec)
	case "SecretText":
		spec, err = json.Marshal(a.SecretTextSpec)
	case "WinRmTGTKeyTabFile":
		spec, err = json.Marshal(a.WinRmTgtKeyTabFileSpec)
	case "WinRmTGTPassword":
		spec, err = json.Marshal(a.WinRmTgtPasswordSpec)
	case "WinRmNTLM":
		spec, err = json.Marshal(a.WinRmNtlmSpec)
	default:
		panic(fmt.Sprintf("unknown secret type %s", a.SecretSpec.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
