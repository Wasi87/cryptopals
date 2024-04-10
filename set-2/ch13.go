package set2

import (
	set1 "cryptopals/set-1"
	"cryptopals/util"
	"fmt"
	"log"
	"strings"
	"sync"
)

var (
	maxUID = 10
	uidMutex sync.Mutex
	once13 sync.Once
	key13 []byte
)

func KvParser(input string) map[string]string {
	strSlice := strings.Split(input, "&")

	kv := make(map[string]string)
	for _, string := range strSlice {
		keyValue := strings.Split(string, "=")
		if len(keyValue) == 2 {
			kv[keyValue[0]] = keyValue[1]
		} else {
			log.Fatal("include invalid strings")
		}
	}
	return kv
}

func getNextUID() int {
	uidMutex.Lock()
	defer uidMutex.Unlock()
	maxUID++
	return maxUID
}

func ProfileFor(email string) map[string]interface{} {
	if strings.ContainsAny(email, "=&") {
		fmt.Println("illegal strings")
	}
	userProfile := map[string]interface{}{
		"email": email,
		"uid": getNextUID(),
		"role": "user",
	}
	return userProfile
}

func ProfileEncode(profile map[string]interface{}) string {
    encoded := fmt.Sprintf(
		"email=%s&uid=%d&role=%s", 
		profile["email"], 
		profile["uid"], 
		profile["role"],
	)

    return encoded
}

func EncryptProfile(profile string, key []byte) []byte {
	profileByte := []byte(profile)
	encProfile := util.AesEncryptECB(key, profileByte)
	return encProfile
}


func Chal13(email string) []byte {
	once13.Do(func(){ key13 = GenAESKey(16)})
	encryptedProfile := EncryptProfile(ProfileEncode(ProfileFor(email)), key13)
	return encryptedProfile
}


func BreakChal13(input []byte) map[string]string {
	decrypted := set1.AesDecryptECB(key13, input)
	return KvParser(string(decrypted))
}
