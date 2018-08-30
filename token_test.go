package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const (
	key = `-----BEGIN RSA PRIVATE KEY-----
keyID: 2PFD:CGJP:D6WJ:3XF3:KFJR:C35Z:STRC:ULJ6:52MX:2R4X:7ACC:FROB

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
-----END RSA PRIVATE KEY-----`
	cert = `-----BEGIN CERTIFICATE-----
MIIFHjCCAwagAwIBAgIBADANBgkqhkiG9w0BAQsFADBGMUQwQgYDVQQDEzsyUEZE
OkNHSlA6RDZXSjozWEYzOktGSlI6QzM1WjpTVFJDOlVMSjY6NTJNWDoyUjRYOjdB
Q0M6RlJPQjAeFw0xODA4MjMyMDU2NDJaFw0yODA4MjcyMDU2NDJaMEYxRDBCBgNV
BAMTOzJQRkQ6Q0dKUDpENldKOjNYRjM6S0ZKUjpDMzVaOlNUUkM6VUxKNjo1Mk1Y
OjJSNFg6N0FDQzpGUk9CMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA
yo9usKZb8grVD3VNdW1s/p1Cogr2Dged6yfM73ysYmPGmbPct80Oz3HIqsUXFGjW
o3gi645crImgonOKse4hW/EGvJDw0qvP7BoiI3pzbFM5yg3Rd9C8sORpuzgxhG/D
G/MmN1ee+JrJHU+Ydk6QSiHPqJglzsLpD7wH/Za5hTucqdbHKUT2YvCyTf+RAD9J
PqEx2U+bHB8e7y/4LgrBJUcpBcA7cIM5rH+JiZ9GG76gXSymzVDNKASVJgKffbDx
82nl54CcdwwXvDIKn9MZ0lE5JHMAFDhAJ+VB7a9Q0Fppq+h91Nap3zpEicl6XrC+
zVn9FFhNBbHBd3C9D5z+RTvoLHy4AC5haTURMCg0Is+I80PZqtC1uhNILaii4Fuy
5K1OOIQXDKm00hf4AcipDdxQgLNu7zxwQxMF4nZB2HM8QPMYsfk85tBRnfN2nuDv
m+bWAZIIbBF/Y9TCaDV//u3lPaArl6cP6S2Kk4GWUnFrfF/8F/KVm85pI//YVrLa
OwdaP1xbrhErMN82R6hAFwHAejvPtN+yRAJUqHWyqfRdZxKtWa6IuHn9fbtU3MSk
8BJJMj+mWFNE6McHnAnolZcLosFYCyoLl28Gu66tqgf9q1t96srTsjW13kikKAan
NumGLBgMiiRVtY+6ZNlX+0iEUXOJZZW1me8oYdoY5n0CAwEAAaMXMBUwEwYDVR0l
BAwwCgYIKwYBBQUHAwIwDQYJKoZIhvcNAQELBQADggIBACGurzxyNcO9LS3saYZG
RiLu8PO47dBqQWjvEN+OwUdLnA+ZfuIDenfNuVtKs7vUjNRpKv2+KnzgCLcdhDJK
9b1FkUqmwdNGVKjsVcsbmqi+vRqzDbKib+WvZ2FQwhsYTHb42zcXp/Dhy2HTmkjH
UMMiWm4gPVvgAvbfi3C0M93YiDEIs49fIDn67CPYn+4eLYiFwR+3sFKXeXXtVdBK
Cdehgc/2YjqFXqCIeBrc6hnMDmuc5K2ECIFoSu8ViVJPSz3SSl6+4armrQFe9T9e
GNiDlaGzrXLNcwNWFPT9u0MQTiRyKhkye8On6dG9xhumuPZevdzBThBmUbehyp3D
+y8K5be9M4QotHbxUTBQgdYzyAlinrKeOOdlaxcauIJ46U7IZHftCsV1BCKh8MVd
1LK7clfLHIhbKKUG0Oa2N6z7O+6R8D87tL98V3tu1IoEOfraQ35aPd0jF0dAaOmp
tJNhVRWhKejlMNQitTEOpZqFMhCgdD6v3uoSyMG2tnd7HDYE9K40jUBfTAN7Iz9k
DCT2hVYvVkkjc9Jsg3VzIQghWak3Elz+WM90HQr1CZZyI3ozRSWislZbJOb+TxsX
dzj9zZb8ipo4oOP9o4t4uT1tJNyLaifV0E+Zo/FDXASnCXcSjOXcLrWDpqbZWMlI
Tjazy+T49bAj/GHz/JU95pxv
-----END CERTIFICATE-----`
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		accesses Accesses
		audience string
		subject  string
		iat      time.Time
		jti      string
	}
	iat := time.Unix(1535663941, 0)
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "empty token",
			args:    args{iat: iat, jti: "PNy9ASCufu44-relzg3L"},
			want:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjJQRkQ6Q0dKUDpENldKOjNYRjM6S0ZKUjpDMzVaOlNUUkM6VUxKNjo1Mk1YOjJSNFg6N0FDQzpGUk9CIn0.eyJpc3MiOiJhdXRoLmRvY2tlci5pbyIsInN1YiI6IiIsImF1ZCI6IiIsImV4cCI6MTUzNTY2NDg0MSwibmJmIjoxNTM1NjYzOTQxLCJpYXQiOjE1MzU2NjM5NDEsImp0aSI6IlBOeTlBU0N1ZnU0NC1yZWx6ZzNMIn0.JF0GmMEOgMn77x0Jl-GBoCqFbzazYq-0VCVoFdo1X94aHx4MlNYcDJbIK4xreIOb4SeEjMdEn47Z9dZUp74i7QPREmXf8F69XxQ86w-TiyVFg15QJLyywHu_40FOrNq4enzx3tRcz9J6E3FlgtgdsYOtxbpEz_5NNvmNIO7sEqjA3_P26-KLiRKShWT1zZFg6UbIl39LoGlvkBnNO-iSIpTaMCnKqyIOeyxAGTGMP3rGJMwQwSxiNTdjW0D3bwcsSARoFH-iNQHdVZ_XgeJ2Bge_PwweDdYLsyE4P0wyg7cOfDWZFWhzmQ2Z3fa4y7cq4N1txyajYS7VbJI-7ct6D8mZ6KwLRYWW3KD9ClhGJrsDPfwmwzZ5qKwVMlYOfF01b0CNBy3FTbHqKrzZ9wNSrRs6_44tqaNOVugMqz0udXQrdPV3KiIV8WUKJwT3VjOYCNHgXVAhNBNeJoI6B0xiN4QNKN8uAHrUlr-p4Q-q7BIEDejWS734XsTMVqzYsOFB-LqdBs4-3qNRE-QLxc4ysW-R1tF_KEPk_ekMaQtEM3BTHo59U8LpIAKpt2hH-AOEKXxmsDIrt8-RDnSRT8WYfKrnzbr0_ZEPY9XWuGIOt5hlGw4tL44FpKglfZ-qzjn_zYr9AIswb9DSarGt_7Dfl8Jpch0mqc4lEfY7W5H5s34",
			wantErr: false,
		},
		{
			name: "empty token",
			args: args{
				accesses: *GetAccesses("reposiotry:test/sample:latest:pull,push repository:cblomart/foo:pull"),
				audience: "registry.docker.io",
				subject:  "cblomart",
				iat:      iat,
				jti:      "IyAUaXE9_5sohIBrORrU",
			},
			want:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjJQRkQ6Q0dKUDpENldKOjNYRjM6S0ZKUjpDMzVaOlNUUkM6VUxKNjo1Mk1YOjJSNFg6N0FDQzpGUk9CIn0.eyJpc3MiOiJhdXRoLmRvY2tlci5pbyIsInN1YiI6ImNibG9tYXJ0IiwiYXVkIjoicmVnaXN0cnkuZG9ja2VyLmlvIiwiZXhwIjoxNTM1NjY0ODQxLCJuYmYiOjE1MzU2NjM5NDEsImlhdCI6MTUzNTY2Mzk0MSwianRpIjoiSXlBVWFYRTlfNXNvaElCck9SclUiLCJhY2Nlc3MiOlt7InR5cGUiOiJyZXBvc2lvdHJ5IiwibmFtZSI6InRlc3Qvc2FtcGxlOmxhdGVzdCIsImFjdGlvbnMiOlsicHVsbCIsInB1c2giXX0seyJ0eXBlIjoicmVwb3NpdG9yeSIsIm5hbWUiOiJjYmxvbWFydC9mb28iLCJhY3Rpb25zIjpbInB1bGwiXX1dfQ.rxBRxWpMnuAR7IYWvTdJbHF0_TYy0OIcu7_hVmup1tcpSBM_jjvEuOp29NxYChn94mTpa1CIGUM7S5fq4yOT2SFBDJksu2YCWWpA8p-fvNs5PqViA4-Nf1-6iKydX-BlVw3Iu7mo5BiyqJKF7d9LIlGjlPkXRG1t-7o0gLTAChI43BDBliAyhDsWG673B4THZ5vAB9AJ9rKBAbZG2h70cBhQPP7kKvlwVUO7AwTa34ZhU5fNgPIQeGjoZBYfvrAFOVoSqY5ai6vfn7emhU41ibjrEFH62_62KHHbgSKhQTgulA6hER8aXsqpPkg2bezie7JzFBHLc38i6fkakl00jT1E8s7ccaN9IZmngnQ2o__vQtIy2enYaXOmxa9jPIygYdRINTkH2oFfkKzDJsf4kI25TU1HxyfFpM8bIFlpWolJnCp-cf97_Asl3DlBxPcsH1o3Ggr0NfU2harZiEV5PeQpmB4Geb2Pcq1uAB-ojhEFGoKjaNfdbxyCNppB-EFUQ46Z6i0Vd2WOfNGCLgpUaaBgmEU654UYs-CqHI4g7JhPTwmNWP5JFgJfs1s8IYE5NpHOtvxPUWKh5UYeEEiD9FMzIX4kgmp19is5xsR6vLMm-UElw9wvlfzilEMdf03s3DSQS1c5tbxmD_KVOl1OsGF9Jub1qA2_Ej9IfNaS2Ck",
			wantErr: false,
		},
	}
	AuthConfig = Config{
		JWSCert: "/tmp/cert.crt",
		JWSKey:  "/tmp/cert.key",
		Issuer:  "auth.docker.io",
	}
	ioutil.WriteFile("/tmp/cert.key", []byte(key), 0600)
	ioutil.WriteFile("/tmp/cert.crt", []byte(cert), 0600)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.accesses, tt.args.audience, tt.args.subject, tt.args.iat, tt.args.jti)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
	os.Remove(AuthConfig.JWSCert)
	os.Remove(AuthConfig.JWSKey)
	AuthConfig = Config{}
}

func TestGenerateJTI(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateJTI(); got != tt.want {
				t.Errorf("GenerateJTI() = %v, want %v", got, tt.want)
			}
		})
	}
}
