package utility

import (
	"bytes"
	"dossier_scolaire/structure"
	"io"
	"mime/multipart"
	"os"
	"strings"

	crypter "github.com/KonishiLee/wechat-crypter"
)

// This function cryped the datas contain in the multipart.File variable and return a formatted string ready to be write in a file
func Crypt(teacher structure.Teacher, in multipart.File) *strings.Reader {
	// We translate the multipart file into a byte buffer to modify it more easily
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, in)

	// We prepare the crypter with the key generate at the creation o the teacher's profile
	msgCrypter, _ := crypter.NewMessageCrypter("", teacher.Key, "")
	msgCrypted, _ := msgCrypter.Encrypt(buf.String())

	// We return the datas crypted and formated to be write in a file
	return strings.NewReader(msgCrypted)
}

// This function cryped a file (who the path is contain in the Teacher struct)
func FullEncrypt(teacher structure.Teacher) {
	// If there is no path we do nothing
	if teacher.SigningUpPath == "" {
		return
	}

	// We open the file
	file, _ := os.ReadFile(teacher.SigningUpPath[1:])

	// We prepare the crypter and crypt the datas of the file
	msgCrypter, _ := crypter.NewMessageCrypter("", teacher.Key, "")
	decrypt, _ := msgCrypter.Encrypt(string(file))

	// We recreate the file
	newFile, _ := os.Create(teacher.SigningUpPath[1:])
	newFile.Write([]byte(decrypt))
}

// This function decryped a file (who the path is contain in the Teacher struct)
func FullDecrypt(teacher structure.Teacher) {
	// If there is no path we do nothing
	if teacher.SigningUpPath == "" {
		return
	}

	// We open the file
	file, _ := os.ReadFile(teacher.SigningUpPath[1:])

	// We prepare the crypter and decrypt the datas of the file
	msgCrypter, _ := crypter.NewMessageCrypter("", teacher.Key, "")
	decrypt, _, _ := msgCrypter.Decrypt(string(file))

	// We recreate the file
	newFile, _ := os.Create(teacher.SigningUpPath[1:])
	newFile.Write(decrypt)
}
