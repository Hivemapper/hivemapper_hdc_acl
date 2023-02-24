package hivemapper_hdc_acl

import (
	"os"
	"path"
	"testing"

	"github.com/streamingfast/solana-go"
	"github.com/stretchr/testify/require"
)

const aclJson = `{"managers":["8phet3DAk2eGdQfvYmtbo2fhFqA5Sr2NFSBCmf3RE737"],"drivers":["2QUvDxdPoRnPaeFbn6QnVAzgJMs7mFQTDc6Q9jc8uHza","2Ryy73QVFoS3deiUfAB5V8zjNCL7UG6DoXC7T5GF91C7","2SfRNgN227dSeK8zQb86CF6MHGy4Uat7JooT3dCUgKj4","2qUoQFTScmhJyMnPT4WB6BakSWufD8sMR1oCFMNxZxDD","32z6LxuaquZvxxDVfuuXE1nUQTZNYkJnFt16XPiXSWWQ","3JFPgUEu8m4a6vfFFYNEKUF4tgxa7hG9FpCsRC7wKXFk","3U78XRD52dpoPVWnTDTtmc4BPS43ST8edf776EGVQZ8t","3mWt5Pz8sorLqpjo66tpSa3c2VNGgtf7aMytJcuMNo4M","4VpfkkNAN9xP8UoKjM3vxtTubj93DDxvE9spJrAJp3ZG","4ZUHYCTbyk8ZjDtsvGebn8tVq61oKbkjBjyExa5243Ts","5FTC5h6K6NbDUantunVM8Tx8kjkfRSdp5yeCD5Kusc9k","5KZwsTp5ExojhSbyxxAzwESCZgcXXxxzzc5Ttz53Z1y4","5SeLXRmhwrYjaE28XeGRqrNP83jJgXPpvDfnA8qEQsRm","73hrbYpEH2iWyPQZKUBvk68JbjauHUEqHusthRi5ca6S","7WyPfL6gZiAu5mon7JcCvA8MPp7gip9oJssqXDVAjJGV","8T9TVEAKRCGpjevus88G2a71NpPU9gADUTWRvs8sqp6v","92EELA1ayt7W8dgmmAiwL2Umb7t57zZdq5iB74kp2kRB","9j3GFWXd6qhKDFT5vrTRYdVnp4hJjKDSLu4eNE73jBjR","AFfBJLWv2FCMzXzzTxwonagnXh8LTVJxw7c1ukCKmEWX","AM7GUBwiwQrtx2QSmmLYWyxxW31RVFzvqeoLY9BYPBLs","ANScf6wTQuMsvJrL8cfSioYw9eFhopacGQjmrxK4yQdz","AhnNiQygBfDMwZ4xQTTkkDcJZ21HdddQtNKucVjNfene","AiwitQFjNbgC7ntRATkwAmZJCTDF4fUfSCiX4TgMVc7L","B65DSMYPCZrVNxM6FnaLo9KQ4t66RYqdDUhwQVZEvVxF","BMiywUStY5s6fwg71tpbsXAqcr8uLcvzw81H4Jx3wBwN","BSF7sWePutQ3fjcVarhRczsCHFnkeTKqqEDjyL3sZVeT","C45HW38AubD5GZfQYqETe8sEf5fEx6ovFA6pvu8VuFB5","C85V9UTBoA5HfS6nizzkoSBPXo58oKNj8C86YeUxh3BX","CGbuvx5255L7igjfrGcjC4QHjyoRwP7aMWPBqgZQeR6F","CN6zcp39HiXRzc7zZQLnaGzSM3xXhA6igMBHSdbTWa2o","CR6CPN8243jJdSSMTnUrmtBv6HT6BtUb2qosazL7boRP","Cnrp1a5unjuNJigrsJxfVn616nZ2tkqQLkpckhuDp2EK","CrmnJ2Txc5pdNJwQm4SvkZbCV9A9fVWaw5Ux5cs2qjXE","DK7HATHN2L6KuTi18Uw88MK4NXkSQiunvThxHjXc5RxW","DeSsnDR8orbsPM5aF4BeogQYBmqF59i29zcsMTyk2gDE","ECLxuppcAqvuignA67WSpH1HS18mzPKJnv3smQCwJ9xf","EGpCigyL3MJbAv4SdwNkjEsFgQt61WYkAQ57ncFoPaNz","ELZX8ofwvRrYGHwFeNiHUMd4gG4kmwEpL5ikQVdaPggN","EbAwLh3zVvP3R1CqHSW15vV3wioPmynDXTk8hSZf9QUS","EvgcuZEqaMSSunbYQ47Tct5ZhHpw8fA7ZxrucGE8E21p","FAcRaPH8NfigW9pKcffpW2vtLGSmYCadJBVu5z62f4tY","FM6vD46KqX7aoEwRiHMfwRbTGDQvhJSGqEpwSenN8sVj","FbqKKyZ4L7VSDECGDEyCWkiiiD8p7dmckQrnMSJXeUoq","FkHGopc8jm9kWwMdbmHeqGY54q7kxeGUBov38e8VAryK","GKJcFpfsxmSZCnLeTKddt1CYQUXFWfTWqerWHV6jaD93","GMupSjCNumVcFStntkMPquziAVCbEjj8yZbdqcsaZ1NX","GV9cGSqXVk12gJKdJUFnDKDYarH1Wzyvt6krxvvVgaqG","H1rZZM9wDYYL1yjNS8khfjqyMvCAmGSpbuVgzPWonxwB","HfpidEwdmxFuirV2MZFx9CD1MNJzqjbo3jsekN7v4ckF","HpuVb2Vd99ALpxGtGqZXyt9aqXTDozfEGYjxCnBt5AEi","Ht5P5F1cKsrsmw57duS8JmiuyH3FjHw31UAyjbnSZijh","Q2vsDnTRtsrBia1jAG7ysvJDhio3DkUmyZiGMyY2mwJ","TGdg99rJowvt9Q86iTebr71fWqjY4nTKH9b1xLLLiXi","iTyL2ZqD2d7ZyCqM8G1rcGRuLLhumGqUZQPE8rN5n2B","sZkrcRGfZwxKDMhnCTjwwpr2JGZT3XVJ3SoPY1Nkc29"]}`
const signatureB58 = `3drYb1k1HmnP6ubDWsVVZx4qsRBhBuWST1YsvTdnDpKhEyTgP1Nrxq53s7mMuEL7MfdQ5uWSZr3yxAsFEFKe7rH2`
const clearAclJson = `{"managers":["3Pa4DNHKyEPJ5YQPaQBRDggstgmd89Zhr4yVMndo6T4C"],"drivers":["AW2MMchomiqbyfKUu1CkUF8n9P41H9y7C6H6MhdYkWXf"]}`
const clearSignatureB58 = `Brz2oAM8YR8m78xLH5CVWgwZdbiCyc2WZ6ZsBYjCtAwjaXkuyXirqueBNSAnowyRFXWvbV1PiyGP2AfBfgiER1u`

const legacyAclJson = `{"managers":["2hBiLi6AQ59knbq8eoonWa4rHS6NdWaqvmA9FBipC5Gf","3bhrVE8tFQYWEcwVgjaXw1FFYx4rorHpMKRuJRGHGjyR","8ZC6P1vjm3WNQ19eCYWWY6SbzWGyF2vaQdyw4DQ5Ky7T","94BBSnkJ2E8SaHRTBVHcLV5ey6EXzPu4oBGKv7ghfKdK","97pv7DTLsDw1AskKQsu5FskcCBpAReVQrLuyyoqxP58q","98juY4BXARPMrwTteeRMEiM51boSFvVHJcyRwsf8niMH","9z2eycrbn6U24qEsJCVYTeqSbbihjRxzFGmifTZhKX7w"],"drivers":["3bhrVE8tFQYWEcwVgjaXw1FFYx4rorHpMKRuJRGHGjyR","3gp5bL9kksMPAgeiV4MEysUBwZ4Bi99s9AksdXdqFsgz","7ZSt2K1SoSDeVwMhsTZ7QAV6TEPXMy2HCySvMouybW6S","8ZC6P1vjm3WNQ19eCYWWY6SbzWGyF2vaQdyw4DQ5Ky7T","97pv7DTLsDw1AskKQsu5FskcCBpAReVQrLuyyoqxP58q","9SNiQuTjkvTrXGDiE3KHfKtSFdzFVHgz9LfvdypqGvYq","9z2eycrbn6U24qEsJCVYTeqSbbihjRxzFGmifTZhKX7w","Cx8kfR2bsL8yxgATkMqxxjtfYFWK1DCASbj2KgADWCSL","GD6mXRyysUWbDmLNkoowhUzQ2gHk6hwSHkYrNXTGx3oF","HvdiwErbqgEJRu6oUSA1PpDodpJ1U3wvhEprtJEajyvM"]}`
const legacySignatureB58 = `4W13pFpEac3V2uoChfubteSQXamffosEXmBnLKLhDb2sa7ufsgKwZUteVDnr4uht83YKsyg2qhwzZwmchoz4RzTF`

func TestAcl_ValidateSignature(t *testing.T) {
	acl, err := NewAclFromData([]byte(aclJson))
	require.NoError(t, err)

	signature, err := solana.NewSignatureFromBase58(signatureB58)
	require.NoError(t, err)

	valid := acl.ValidateStoreSignature(signature)
	require.True(t, valid)
}

func TestAcl_EmptyAcl(t *testing.T) {
	aclFolder := "/tmp/acl"
	os.RemoveAll(aclFolder)

	var acl *Acl
	acl = &Acl{}
	//acl, err := NewAclFromData([]byte(""))
	//require.NoError(t, err)
	//fmt.Println(acl)
	signature, err := solana.NewSignatureFromBase58(signatureB58)
	require.NoError(t, err)

	acl.Store(aclFolder, signature)
	//require.NoError(t, err)
	//
	//valid := acl.ValidateStoreSignature(signature)
	//require.True(t, valid)
}

func TestAcl_ValidateLegacySignature(t *testing.T) {
	acl, err := NewAclFromData([]byte(legacyAclJson))
	require.NoError(t, err)

	signature, err := solana.NewSignatureFromBase58(legacySignatureB58)
	require.NoError(t, err)

	valid := acl.ValidateStoreSignature(signature)
	require.True(t, valid)
}

func Test_messageToSign(t *testing.T) {
	acl, err := NewAclFromData([]byte(aclJson))
	require.NoError(t, err)

	message, err := acl.storeMessageToSign()
	require.NoError(t, err)

	require.Equal(t, "Access Control List with 1 manager(s) and 55 driver(s). Hash: 492238a930883b0dd6ed791dc1e0152a", string(message))
}

func Test_messageClearToSign(t *testing.T) {
	acl, err := NewAclFromData([]byte(aclJson))
	acl.FleetName = "Pere Noel Hivrogne"
	require.NoError(t, err)

	message, err := acl.clearMessageToSign()
	require.NoError(t, err)

	require.Equal(t, "Clearing Access Control List for fleet Pere Noel Hivrogne", string(message))
}

func Test_ValidateLegacyMessageClearToSign(t *testing.T) {
	acl, err := NewAclFromData([]byte(clearAclJson))
	acl.FleetName = "graveful-thistle-pig"
	signature := "2TD922GcdnkkVedKeZoy9XGPzGkF7Z9erQt1z3ktGgHCUYVcxbMYXx9wwvUHqF9P2K9hGs9D7j9pngCTLQeup8gx"
	signatureBS58, err := solana.NewSignatureFromBase58(signature)
	_ = signature
	require.NoError(t, err)

	message, err := acl.clearMessageToSign()
	require.NoError(t, err)

	require.Equal(t, "Clearing Access Control List for fleet graveful-thistle-pig", string(message))
	require.True(t, acl.validateSignature(message, signatureBS58))
}

func Test_AclClearFromDevice(t *testing.T) {
	aclFolder := "/tmp/acl"
	os.RemoveAll(aclFolder)

	acl, err := NewAclFromData([]byte(aclJson))
	signature, err := solana.NewSignatureFromBase58(signatureB58)

	err = acl.Store(aclFolder, signature)
	require.NoError(t, err)

	exist := AclExistOnDevice(aclFolder)
	require.True(t, exist)

	_, err = os.Stat(path.Join(aclFolder, AclFileName))
	require.NoError(t, err)

	err = AclClearFromDevice(aclFolder, "")
	require.NoError(t, err)

	_, err = os.Stat(path.Join(aclFolder, AclFileName))
	require.True(t, os.IsNotExist(err))
	_, err = os.Stat(path.Join(aclFolder, AclSignatureFileName))
	require.True(t, os.IsNotExist(err))
}

func Test_AclClearFromDeviceWithSign(t *testing.T) {
	aclFolder := "/tmp/acl"
	os.RemoveAll(aclFolder)

	acl, err := NewAclFromData([]byte(clearAclJson))
	acl.Version = "9.9.9"
	acl.FleetName = "Pere Noel Hivrogne"
	signature, err := solana.NewSignatureFromBase58(clearSignatureB58)

	err = acl.Store(aclFolder, signature)
	require.NoError(t, err)

	_, err = os.Stat(path.Join(aclFolder, AclFileName))
	require.NoError(t, err)

	err = AclClearFromDevice(aclFolder, "")
	require.ErrorIs(t, ErrSignatureRequired, err)

	err = AclClearFromDevice(aclFolder, signature.String())
	require.NoError(t, err)

	_, err = os.Stat(path.Join(aclFolder, AclFileName))
	require.True(t, os.IsNotExist(err))
	_, err = os.Stat(path.Join(aclFolder, AclSignatureFileName))
	require.True(t, os.IsNotExist(err))
}
