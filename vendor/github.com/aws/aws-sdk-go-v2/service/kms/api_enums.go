// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

type AlgorithmSpec string

// Enum values for AlgorithmSpec
const (
	AlgorithmSpecRsaesPkcs1V15   AlgorithmSpec = "RSAES_PKCS1_V1_5"
	AlgorithmSpecRsaesOaepSha1   AlgorithmSpec = "RSAES_OAEP_SHA_1"
	AlgorithmSpecRsaesOaepSha256 AlgorithmSpec = "RSAES_OAEP_SHA_256"
)

func (enum AlgorithmSpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AlgorithmSpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConnectionErrorCodeType string

// Enum values for ConnectionErrorCodeType
const (
	ConnectionErrorCodeTypeInvalidCredentials       ConnectionErrorCodeType = "INVALID_CREDENTIALS"
	ConnectionErrorCodeTypeClusterNotFound          ConnectionErrorCodeType = "CLUSTER_NOT_FOUND"
	ConnectionErrorCodeTypeNetworkErrors            ConnectionErrorCodeType = "NETWORK_ERRORS"
	ConnectionErrorCodeTypeInternalError            ConnectionErrorCodeType = "INTERNAL_ERROR"
	ConnectionErrorCodeTypeInsufficientCloudhsmHsms ConnectionErrorCodeType = "INSUFFICIENT_CLOUDHSM_HSMS"
	ConnectionErrorCodeTypeUserLockedOut            ConnectionErrorCodeType = "USER_LOCKED_OUT"
	ConnectionErrorCodeTypeUserNotFound             ConnectionErrorCodeType = "USER_NOT_FOUND"
	ConnectionErrorCodeTypeUserLoggedIn             ConnectionErrorCodeType = "USER_LOGGED_IN"
)

func (enum ConnectionErrorCodeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConnectionErrorCodeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConnectionStateType string

// Enum values for ConnectionStateType
const (
	ConnectionStateTypeConnected     ConnectionStateType = "CONNECTED"
	ConnectionStateTypeConnecting    ConnectionStateType = "CONNECTING"
	ConnectionStateTypeFailed        ConnectionStateType = "FAILED"
	ConnectionStateTypeDisconnected  ConnectionStateType = "DISCONNECTED"
	ConnectionStateTypeDisconnecting ConnectionStateType = "DISCONNECTING"
)

func (enum ConnectionStateType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConnectionStateType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CustomerMasterKeySpec string

// Enum values for CustomerMasterKeySpec
const (
	CustomerMasterKeySpecRsa2048          CustomerMasterKeySpec = "RSA_2048"
	CustomerMasterKeySpecRsa3072          CustomerMasterKeySpec = "RSA_3072"
	CustomerMasterKeySpecRsa4096          CustomerMasterKeySpec = "RSA_4096"
	CustomerMasterKeySpecEccNistP256      CustomerMasterKeySpec = "ECC_NIST_P256"
	CustomerMasterKeySpecEccNistP384      CustomerMasterKeySpec = "ECC_NIST_P384"
	CustomerMasterKeySpecEccNistP521      CustomerMasterKeySpec = "ECC_NIST_P521"
	CustomerMasterKeySpecEccSecgP256k1    CustomerMasterKeySpec = "ECC_SECG_P256K1"
	CustomerMasterKeySpecSymmetricDefault CustomerMasterKeySpec = "SYMMETRIC_DEFAULT"
)

func (enum CustomerMasterKeySpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CustomerMasterKeySpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DataKeyPairSpec string

// Enum values for DataKeyPairSpec
const (
	DataKeyPairSpecRsa2048       DataKeyPairSpec = "RSA_2048"
	DataKeyPairSpecRsa3072       DataKeyPairSpec = "RSA_3072"
	DataKeyPairSpecRsa4096       DataKeyPairSpec = "RSA_4096"
	DataKeyPairSpecEccNistP256   DataKeyPairSpec = "ECC_NIST_P256"
	DataKeyPairSpecEccNistP384   DataKeyPairSpec = "ECC_NIST_P384"
	DataKeyPairSpecEccNistP521   DataKeyPairSpec = "ECC_NIST_P521"
	DataKeyPairSpecEccSecgP256k1 DataKeyPairSpec = "ECC_SECG_P256K1"
)

func (enum DataKeyPairSpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DataKeyPairSpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DataKeySpec string

// Enum values for DataKeySpec
const (
	DataKeySpecAes256 DataKeySpec = "AES_256"
	DataKeySpecAes128 DataKeySpec = "AES_128"
)

func (enum DataKeySpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DataKeySpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EncryptionAlgorithmSpec string

// Enum values for EncryptionAlgorithmSpec
const (
	EncryptionAlgorithmSpecSymmetricDefault EncryptionAlgorithmSpec = "SYMMETRIC_DEFAULT"
	EncryptionAlgorithmSpecRsaesOaepSha1    EncryptionAlgorithmSpec = "RSAES_OAEP_SHA_1"
	EncryptionAlgorithmSpecRsaesOaepSha256  EncryptionAlgorithmSpec = "RSAES_OAEP_SHA_256"
)

func (enum EncryptionAlgorithmSpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EncryptionAlgorithmSpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExpirationModelType string

// Enum values for ExpirationModelType
const (
	ExpirationModelTypeKeyMaterialExpires       ExpirationModelType = "KEY_MATERIAL_EXPIRES"
	ExpirationModelTypeKeyMaterialDoesNotExpire ExpirationModelType = "KEY_MATERIAL_DOES_NOT_EXPIRE"
)

func (enum ExpirationModelType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExpirationModelType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type GrantOperation string

// Enum values for GrantOperation
const (
	GrantOperationDecrypt                             GrantOperation = "Decrypt"
	GrantOperationEncrypt                             GrantOperation = "Encrypt"
	GrantOperationGenerateDataKey                     GrantOperation = "GenerateDataKey"
	GrantOperationGenerateDataKeyWithoutPlaintext     GrantOperation = "GenerateDataKeyWithoutPlaintext"
	GrantOperationReEncryptFrom                       GrantOperation = "ReEncryptFrom"
	GrantOperationReEncryptTo                         GrantOperation = "ReEncryptTo"
	GrantOperationSign                                GrantOperation = "Sign"
	GrantOperationVerify                              GrantOperation = "Verify"
	GrantOperationGetPublicKey                        GrantOperation = "GetPublicKey"
	GrantOperationCreateGrant                         GrantOperation = "CreateGrant"
	GrantOperationRetireGrant                         GrantOperation = "RetireGrant"
	GrantOperationDescribeKey                         GrantOperation = "DescribeKey"
	GrantOperationGenerateDataKeyPair                 GrantOperation = "GenerateDataKeyPair"
	GrantOperationGenerateDataKeyPairWithoutPlaintext GrantOperation = "GenerateDataKeyPairWithoutPlaintext"
)

func (enum GrantOperation) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GrantOperation) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyManagerType string

// Enum values for KeyManagerType
const (
	KeyManagerTypeAws      KeyManagerType = "AWS"
	KeyManagerTypeCustomer KeyManagerType = "CUSTOMER"
)

func (enum KeyManagerType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyManagerType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyState string

// Enum values for KeyState
const (
	KeyStateEnabled         KeyState = "Enabled"
	KeyStateDisabled        KeyState = "Disabled"
	KeyStatePendingDeletion KeyState = "PendingDeletion"
	KeyStatePendingImport   KeyState = "PendingImport"
	KeyStateUnavailable     KeyState = "Unavailable"
)

func (enum KeyState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyUsageType string

// Enum values for KeyUsageType
const (
	KeyUsageTypeSignVerify     KeyUsageType = "SIGN_VERIFY"
	KeyUsageTypeEncryptDecrypt KeyUsageType = "ENCRYPT_DECRYPT"
)

func (enum KeyUsageType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyUsageType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MessageType string

// Enum values for MessageType
const (
	MessageTypeRaw    MessageType = "RAW"
	MessageTypeDigest MessageType = "DIGEST"
)

func (enum MessageType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MessageType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OriginType string

// Enum values for OriginType
const (
	OriginTypeAwsKms      OriginType = "AWS_KMS"
	OriginTypeExternal    OriginType = "EXTERNAL"
	OriginTypeAwsCloudhsm OriginType = "AWS_CLOUDHSM"
)

func (enum OriginType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OriginType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SigningAlgorithmSpec string

// Enum values for SigningAlgorithmSpec
const (
	SigningAlgorithmSpecRsassaPssSha256      SigningAlgorithmSpec = "RSASSA_PSS_SHA_256"
	SigningAlgorithmSpecRsassaPssSha384      SigningAlgorithmSpec = "RSASSA_PSS_SHA_384"
	SigningAlgorithmSpecRsassaPssSha512      SigningAlgorithmSpec = "RSASSA_PSS_SHA_512"
	SigningAlgorithmSpecRsassaPkcs1V15Sha256 SigningAlgorithmSpec = "RSASSA_PKCS1_V1_5_SHA_256"
	SigningAlgorithmSpecRsassaPkcs1V15Sha384 SigningAlgorithmSpec = "RSASSA_PKCS1_V1_5_SHA_384"
	SigningAlgorithmSpecRsassaPkcs1V15Sha512 SigningAlgorithmSpec = "RSASSA_PKCS1_V1_5_SHA_512"
	SigningAlgorithmSpecEcdsaSha256          SigningAlgorithmSpec = "ECDSA_SHA_256"
	SigningAlgorithmSpecEcdsaSha384          SigningAlgorithmSpec = "ECDSA_SHA_384"
	SigningAlgorithmSpecEcdsaSha512          SigningAlgorithmSpec = "ECDSA_SHA_512"
)

func (enum SigningAlgorithmSpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SigningAlgorithmSpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WrappingKeySpec string

// Enum values for WrappingKeySpec
const (
	WrappingKeySpecRsa2048 WrappingKeySpec = "RSA_2048"
)

func (enum WrappingKeySpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WrappingKeySpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
