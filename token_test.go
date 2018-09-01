package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const (
	key = `-----BEGIN RSA PRIVATE KEY-----
keyID: M4LN:WJ67:67UM:P5UI:ZVOZ:6HKD:DBU7:NUUU:2CMU:YGGF:PAMR:W72Y

MIIEpgIBAAKCAQEA1KGjPmkYUqsTFOgeLQiBbmsNBeGTx7MJ/aLpP7CqLJpMsZ6h
c4Yd3+MzPaHhcLEuQukVSYtoCESirfcu4Em6cAfItzaMcLdPQH1NEqofLfnO/zZB
GckktjvNePywwvvaHB1H1oQJ/EXIOSqktASmVUpgQ1Q+j/Q7wQxPAF2HA1nxDSiX
1PpNLtDcaXvdbuK//j83w0nL2Zt2xsPInREvIvhLV4Ydmi+O0dmu1j8ggzXzZqf8
Y4FGlbx6OPV67TxPO+OK00MIv9Q/dsD5BNFV021hAag9iryMVYYd2+ACU1/hx+cF
KLvVlpQHtrFsV+/AOhIVxK+r+hjstCfavtZjawIDAQABAoIBAQC14j49mLCRpn0A
GT4Wz3vH9BKDwV4kKOaW69ASBxDKboLEPSlmJIdRiIvAYMTzHNyxp6fp7THkQLHX
leWnPeoZAs+SyTqBAIzuYUCYgqWBMnr6nHI7mG+q4qT0pVKet8ju260jtnbEMP0A
yZVx17hgpOqLLXkGsqiajejWKyrelXUNQ4Ros0WPEcpOKvCc5w3BXRG8x9wX5fTL
e8ZzZ1+FHumdVDlWV3BUHXE0fIuorVUy/Q2kJxJHXSpBqtadVwAk6WxP0Re+LTne
lkQeyZ9KFKAgXDpAToSl28ehllsOMKPiucdxeSuHbC2BFBlSEqI4AlyyK3JmlizK
9P27IHphAoGBAPqZa8df74DLdGLz5E1Hl62n2+TzMNq3s7FuOgqy28uY3y3KLxDR
FBscaswUJzBjjAwEWkTdR4L7ajrAD07uHCTDcHaTjYitRurscByl8kS4aOwPJj4J
RWtvcB8pYbYyHayOSp/Ji8fapFcRHvao0PpJiyVpNljXZ2j/7uf+LSy9AoGBANk2
voq7/UVorf1uszSRWRaADvTOjkeWUwt2h0kNRW5xLAbT8tRnj/79d3usa/+ENtro
IRusDTOsOYIYHr2mZk0V0s5uFI8ukWKWewG4VLpTDu53qpd+8Ow8Wt6m9VW8RDnZ
ae9bkXZgzs6GQbGyWoCNfESr4X50Gl1uZq78EhdHAoGBAI6elhpH3RSdtcVDLKFW
p1oreBga16kCd3/5TYsdM73xaMK0mIinlRvA1n8C0TLW1LNaHs2mabZ/w3tgJtYy
6U169Rxm6Vvp0byNh1imkPTPRtdh38/akumL6HGzqWp5py23ZXz+xVCefH0Yv1h9
x0FKbaiC8m0pWi8FyT+7Bpa5AoGBAJE3Pi+1+l762sdByODuAyc9ucIEja5iW2ag
eVVnX/G0C4ENFZzXF7ebcNPns9QBakLzSZ8caT8Qdun4giS8KEyEuIh1o50Nvviw
LdA6kbp3aNYYfp4Fqb/locKU0BPfZ6VdKqtxBlCj6966dxT7bfHfpSKr3ncR28Z2
1oNJ1jZlAoGBAKAp2wIvSko/9Gpd9t3lXuxQz5YGBsLZIE/I+g5DrrG3pzws4B2G
D3sMsdZWsnOSd48E1GogSZKeL7zK2oDBsQlqK0hfj6n3xCCyANpeAuurjQ/EhzOF
9H1l9kWeWvJRQIVUZlHo+t92zIhq+PrZK8AEOwhPBYTSaiLs7L1PRLPn
-----END RSA PRIVATE KEY-----`
	cert = `-----BEGIN CERTIFICATE-----
MIIDHjCCAgagAwIBAgIBADANBgkqhkiG9w0BAQsFADBGMUQwQgYDVQQDEztNNExO
OldKNjc6NjdVTTpQNVVJOlpWT1o6NkhLRDpEQlU3Ok5VVVU6MkNNVTpZR0dGOlBB
TVI6VzcyWTAeFw0xODA4MjMyMjMzMzVaFw0yODA4MjcyMjMzMzVaMEYxRDBCBgNV
BAMTO000TE46V0o2Nzo2N1VNOlA1VUk6WlZPWjo2SEtEOkRCVTc6TlVVVToyQ01V
OllHR0Y6UEFNUjpXNzJZMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
1KGjPmkYUqsTFOgeLQiBbmsNBeGTx7MJ/aLpP7CqLJpMsZ6hc4Yd3+MzPaHhcLEu
QukVSYtoCESirfcu4Em6cAfItzaMcLdPQH1NEqofLfnO/zZBGckktjvNePywwvva
HB1H1oQJ/EXIOSqktASmVUpgQ1Q+j/Q7wQxPAF2HA1nxDSiX1PpNLtDcaXvdbuK/
/j83w0nL2Zt2xsPInREvIvhLV4Ydmi+O0dmu1j8ggzXzZqf8Y4FGlbx6OPV67TxP
O+OK00MIv9Q/dsD5BNFV021hAag9iryMVYYd2+ACU1/hx+cFKLvVlpQHtrFsV+/A
OhIVxK+r+hjstCfavtZjawIDAQABoxcwFTATBgNVHSUEDDAKBggrBgEFBQcDAjAN
BgkqhkiG9w0BAQsFAAOCAQEAkQkasgHmDm5pt2k0O7h+uZWCTl7DqAELhR6QbnY/
WIwIMLqnbYiYyJs5bk+WMY79CE2aT/z2jhgdw4EqUF7Lk62sKGyqpxu8QLO1p308
mN044rGOohb5fSrF5Hw7EpS9KdHlAwpxYzSwlJ7TyGU6fme/spisLg5pDIyKO1b3
2bT7n4+/D0z6Hz5N7N/7QNWvbTT8zOxasz/WW3ZvyVCgM/u+DnGP+Pcwx5pK1Cy/
dac7WZrPg4yVymdGxtqyR+D8GF/rK+9jUBUkTvigma/htBSawyCyMwwoHiTZRX8t
HwIozQVzYxsPicPBWfnGosshqvO6/0cbxGrCQjREc2GyaQ==
-----END CERTIFICATE-----`
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		accesses Scopes
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
			want:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik00TE46V0o2Nzo2N1VNOlA1VUk6WlZPWjo2SEtEOkRCVTc6TlVVVToyQ01VOllHR0Y6UEFNUjpXNzJZIn0.eyJpc3MiOiJhdXRoLmRvY2tlci5pbyIsInN1YiI6IiIsImF1ZCI6IiIsImV4cCI6MTUzNTY2NDg0MSwibmJmIjoxNTM1NjYzOTQxLCJpYXQiOjE1MzU2NjM5NDEsImp0aSI6IlBOeTlBU0N1ZnU0NC1yZWx6ZzNMIn0.iL-kGQfsNHxUYnTzEKymK-EboZAAPy_nVen7Rnmndjcfan4SEKapaegx67-ByfgMU6EqnOY9UeXcjcfF3bB0szwoKvGwLm2LTSmb3Qjenkr5pOiRRQtk0gxGvr0AzheckGYTZa6Qpbzn7pDjXNxU8fLpmHsjvsbLTJb0PemnXmBqSrBAPc80U-IuGdtTkDTqaj9_po4rsWhqKmD3Y0SEoRvrr48NeA6mzcw6Nk-CCaJK2pkLzPRMqdltC1bvYi3U7SmgnlxuELznea7L_GWbU_5mQvoCy_J7pyJEtAmqdjyFdtCZt71PpC57G1bEOQAOqlQ28xDVi9yRk-d9vu02Pg",
			wantErr: false,
		},
		{
			name: "empty token",
			args: args{
				accesses: *GetScopes("reposiotry:test/sample:latest:pull,push repository:cblomart/foo:pull"),
				audience: "registry.docker.io",
				subject:  "cblomart",
				iat:      iat,
				jti:      "IyAUaXE9_5sohIBrORrU",
			},
			want:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik00TE46V0o2Nzo2N1VNOlA1VUk6WlZPWjo2SEtEOkRCVTc6TlVVVToyQ01VOllHR0Y6UEFNUjpXNzJZIn0.eyJpc3MiOiJhdXRoLmRvY2tlci5pbyIsInN1YiI6ImNibG9tYXJ0IiwiYXVkIjoicmVnaXN0cnkuZG9ja2VyLmlvIiwiZXhwIjoxNTM1NjY0ODQxLCJuYmYiOjE1MzU2NjM5NDEsImlhdCI6MTUzNTY2Mzk0MSwianRpIjoiSXlBVWFYRTlfNXNvaElCck9SclUiLCJhY2Nlc3MiOlt7InR5cGUiOiJyZXBvc2lvdHJ5IiwibmFtZSI6InRlc3Qvc2FtcGxlOmxhdGVzdCIsImFjdGlvbnMiOlsicHVsbCIsInB1c2giXX0seyJ0eXBlIjoicmVwb3NpdG9yeSIsIm5hbWUiOiJjYmxvbWFydC9mb28iLCJhY3Rpb25zIjpbInB1bGwiXX1dfQ.ECn_nyH9qI0_l-_tDrN_zQQ1OqJS-QxSftJVnt0YhINqsq8Geol2u1RGu_xxl6FxNdAQ_9CvAoaZLl0nsBSg4e0aD8nLQrEM-wJtEMgkgSvDegltjYd87_mYHoS1-Hse-pzlwO3C-RIKqGoLSfYZk7VKLLFopX0M6J4edC2EWKuklxBQ7i-MQjrQpSVtBwGOlk4m9aI3nVhjXjiXfoxMqztosEhqDQ6CooeXvTLOlLyfkQu57dhwa83PylU3Lj6knxLzgm-IdY4bngZuzhICa8UcQO2cF7bKkVOHV0qkyG7srQ1s5_HPXFM933q8FZknhI4Zk36CkrnbruuUNTXpYw",
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
	}{
		{
			name: "JsonToken ID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GenerateJTI(); len(got) == 0 {
				t.Errorf("GenerateJTI() = empty")
			}
		})
	}
}
