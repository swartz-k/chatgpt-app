//package auth
//
//import (
//	"crypto/tls"
//	"github.com/swartz-k/chatgpt-app/pkg/log"
//)
//
//type Conn struct {
//	Email string
//	Password string
//	Proxy string
//	SessionToken string
//	AccessToken string
//	Session tls.ClientSessionCache
//	// for connect
//	endpoint string
//	headers map[string]string
//}
//
//func (a *Conn) Begin() {
//	log.V(100).Info("Beginning auth process")
//	if a.Proxy != "" {
//
//	}
//}
//
//func NewConn() *Conn {
//	return &Conn{
//		endpoint: "https://chat.openai.com/auth/login",
//		headers: map[string]string{
//			"Host":       "ask.openai.com",
//			"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
//			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
//			"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
//			"Accept-Encoding": "gzip, deflate, br",
//			"Connection":      "keep-alive",
//		},
//	}
//}
//
//response = self.session.get(url=url, headers=headers)
//if response.status_code == 200:
//self.part_two()
//else:
//self.debugger.log("Error in part one")
//self.debugger.log("Response: ", end="")
//self.debugger.log(response.text)
//self.debugger.log("Status code: ", end="")
//self.debugger.log(response.status_code)
//raise Exception("API error")
//
//func (c *Conn) csrf() {
//	log.V(100).Info("Beginning CSRF")
//	url := "https://chat.openai.com/api/auth/csrf"
//	headers := {
//		"Host": "ask.openai.com",
//		"Accept": "*/*",
//		"Connection": "keep-alive",
//		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15 Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
//		"Referer": "https://chat.openai.com/auth/login",
//		"Accept-Encoding": "gzip, deflate, br",
//	}
//}
//response = self.session.get(url=url, headers=headers)
//if response.status_code == 200 and "json" in response.headers["Content-Type"]:
//csrf_token = response.json()["csrfToken"]
//self.part_three(token=csrf_token)
//
//func (c *Conn) signin() {
//	log.V(100).Info("Beginning sign in")
//	url := "https://chat.openai.com/api/auth/signin/auth0?prompt=login"
//
//	payload := "callbackUrl=%2F&csrfToken={token}&json=true"
//	headers := map[string]string{
//		"Host": "ask.openai.com",
//		"Origin": "https://chat.openai.com",
//		"Connection": "keep-alive",
//		"Accept": "*/*",
//		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
//		"Referer": "https://chat.openai.com/auth/login",
//		"Content-Length": "100",
//		"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
//		"Content-Type": "application/x-www-form-urlencoded",
//	}
//	response = self.session.post(url = url, headers = headers, data=payload)
//	if response.status_code == 200 and
//	"json"
//	in
//	response.headers["Content-Type"]:
//	url = response.json()["url"]
//	if url == "https://chat.openai.com/api/auth/error?error=OAuthSignin" or
//	'error'
//	in
//url:
//	self.debugger.log("You have been rate limited")
//	raise
//	Exception("You have been rate limited.")
//	self.part_four(url = url)
//	elif
//	response.status_code == 400:
//	self.debugger.log("Error in part three")
//	self.debugger.log("Invalid credentials")
//	raise
//	Exception("Invalid credentials")
//	else:
//	self.debugger.log("Error in part three")
//	self.debugger.log("Response: ", end = "")
//	self.debugger.log(response.text)
//	self.debugger.log("Status code: ", end = "")
//	self.debugger.log(response.status_code)
//	raise
//	Exception("Unknown error")
//
//}
//
//func (c *Conn) oauth(url string) {
//	log.V(100).Info("Beginning oauth signin")
//	headers := map[string]string{
//		"Host":            "auth0.openai.com",
//		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
//		"Connection":      "keep-alive",
//		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
//		"Accept-Language": "en-US,en;q=0.9",
//		"Referer":         "https://chat.openai.com/",
//	}
//
//	response = self.session.get(url = url, headers = headers)
//	if response.status_code == 302:
//try:
//	state = re.findall(r
//	"state=(.*)", response.text)[0]
//state = state.split('"')[0]
//self.part_five(state= state)
//except IndexError:
//self.debugger.log("Error in part four")
//self.debugger.log("Status code: ", end = "")
//self.debugger.log(response.status_code)
//self.debugger.log("Rate limit hit")
//self.debugger.log("Response: " + str(response.text))
//raise Exception("Rate limit hit")
//else:
//self.debugger.log("Error in part four")
//self.debugger.log("Response: ", end = "")
//self.debugger.log(response.text)
//self.debugger.log("Status code: ", end = "")
//self.debugger.log(response.status_code)
//self.debugger.log("Wrong response code")
//raise Exception("Unknown error")
//
//}
//
//func (c *Conn) identifier() {
//	log.V(100).Info("Beginning identifier")
//	url := "https://auth0.openai.com/u/login/identifier?state={state}"
//
//	headers := map[string]string{
//		"Host":            "auth0.openai.com",
//		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
//		"Connection":      "keep-alive",
//		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
//		"Accept-Language": "en-US,en;q=0.9",
//		"Referer":         "https://chat.openai.com/",
//	}
//
//	response = self.session.get(url, headers = headers)
//	if response.status_code == 200:
//	if re.search(r'<img[^>]+alt="captcha"[^>]+>', response.text):
//	self.debugger.log("Error in part five")
//	self.debugger.log("Captcha detected")
//	raise
//	ValueError("Captcha detected")
//	self.part_six(state = state, captcha = None) else:
//	self.debugger.log("Error in part five")
//	self.debugger.log("Response: ", end = "")
//	self.debugger.log(response.text)
//	self.debugger.log("Status code: ", end = "")
//	self.debugger.log(response.status_code)
//	raise
//	ValueError("Invalid response code")
//
//}
//
//func (c *Conn) passIdentifier( state, captcha string) {
//	log.V(100).Info("Beginning pass identifier")
//	url := "https://auth0.openai.com/u/login/identifier?state={state}"
//	email_url_encoded = url_encode(self.email_address)
//	payload := "state={state}&username={email_url_encoded}&captcha={captcha}&js-available=true&webauthn-available=true&is-brave=false&webauthn-platform-available=true&action=default "
//
//	if captcha is
//None:
//	payload = (
//		f
//	"state={state}&username={email_url_encoded}&js-available=false&webauthn-available=true&is"
//	f
//	"-brave=false&webauthn-platform-available=true&action=default "
//	)
//
//	headers =
//	{
//		"Host": "auth0.openai.com",
//		"Origin": "https://auth0.openai.com",
//		"Connection": "keep-alive",
//		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
//		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) "
//		"Version/16.1 Safari/605.1.15",
//			"Referer": f
//		"https://auth0.openai.com/u/login/identifier?state={state}",
//			"Accept-Language": "en-US,en;q=0.9",
//		"Content-Type": "application/x-www-form-urlencoded",
//	}
//	response = self.session.post(url, headers = headers, data = payload)
//	if response.status_code == 302:
//	self.part_seven(state = state) else:
//	self.debugger.log("Error in part six")
//	self.debugger.log("Response: ", end = "")
//	self.debugger.log(response.text)
//	self.debugger.log("Status code: ", end = "")
//	self.debugger.log(response.status_code)
//	raise
//	Exception("Unknown error")
//}
//
//
//def part_seven(self, state: str) -> None:
//self.debugger.log("Beginning part seven")
//"""
//We enter the password
//:param state:
//:return:
//"""
//url = f"https://auth0.openai.com/u/login/password?state={state}"
//
//email_url_encoded = self.url_encode(self.email_address)
//password_url_encoded = self.url_encode(self.password)
//payload = f"state={state}&username={email_url_encoded}&password={password_url_encoded}&action=default"
//headers = {
//"Host": "auth0.openai.com",
//"Origin": "https://auth0.openai.com",
//"Connection": "keep-alive",
//"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
//"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) "
//"Version/16.1 Safari/605.1.15",
//"Referer": f"https://auth0.openai.com/u/login/password?state={state}",
//"Accept-Language": "en-US,en;q=0.9",
//"Content-Type": "application/x-www-form-urlencoded",
//}
//try:
//response = self.session.post(url, headers=headers, data=payload)
//self.debugger.log("Request went through")
//except Exception as e:
//self.debugger.log("Error in part seven")
//self.debugger.log("Exception: ", end="")
//self.debugger.log(e)
//raise Exception("Could not get response")
//if response.status_code == 302:
//self.debugger.log("Response code is 302")
//try:
//new_state = re.findall(r"state=(.*)", response.text)[0]
//new_state = new_state.split('"')[0]
//self.debugger.log("New state found")
//self.part_eight(old_state=state, new_state=new_state)
//except Exception as e:
//self.debugger.log("Error in part seven")
//self.debugger.log("Exception: ", end="")
//self.debugger.log(e)
//raise Exception("State not found")
//elif response.status_code == 400:
//self.debugger.log("Error in part seven")
//self.debugger.log("Status code: ", end="")
//self.debugger.log(response.status_code)
//raise Exception("Wrong email or password")
//else:
//self.debugger.log("Error in part seven")
//self.debugger.log("Status code: ", end="")
//self.debugger.log(response.status_code)
//raise Exception("Wrong status code")
//
//def part_eight(self, old_state: str, new_state) -> None:
//self.debugger.log("Beginning part eight")
//url = f"https://auth0.openai.com/authorize/resume?state={new_state}"
//headers = {
//"Host": "auth0.openai.com",
//"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
//"Connection": "keep-alive",
//"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) "
//"Version/16.1 Safari/605.1.15",
//"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
//"Referer": f"https://auth0.openai.com/u/login/password?state={old_state}",
//}
//response = self.session.get(url, headers=headers, allow_redirects=True)
//is_200 = response.status_code == 200
//if is_200:
//# Access Token
//access_token = re.findall(
//r"accessToken\":\"(.*)\"",
//response.text,
//)
//if access_token:
//try:
//access_token = access_token[0]
//access_token = access_token.split('"')[0]
//except Exception as e:
//self.debugger.log("Error in part eight")
//self.debugger.log("Response: ", end="")
//self.debugger.log(response.text)
//self.debugger.log("Status code: ", end="")
//self.debugger.log(response.status_code)
//raise e
//else:
//self.debugger.log("Error in part eight")
//self.debugger.log("Response: ", end="")
//self.debugger.log(response.text)
//self.debugger.log("Status code: ", end="")
//self.debugger.log(response.status_code)
//raise Exception("Auth0 did not issue an access token")
//self.part_nine()
//else:
//self.debugger.log("Incorrect response code in part eight")
//raise Exception("Incorrect response code")
//
//def save_access_token(self, access_token: str) -> None:
//"""
//Save access_token and an hour from now on ./Classes/auth.json
//:param access_token:
//:return:
//"""
//self.access_token = access_token
//
//def part_nine(self) -> bool:
//self.debugger.log("Beginning part nine")
//url = "https://chat.openai.com/api/auth/session"
//headers = {
//"Host": "ask.openai.com",
//"Connection": "keep-alive",
//"If-None-Match": '"bwc9mymkdm2"',
//"Accept": "*/*",
//"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) "
//"Version/16.1 Safari/605.1.15",
//"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
//"Referer": "https://chat.openai.com/chat",
//"Accept-Encoding": "gzip, deflate, br",
//}
//response = self.session.get(url, headers=headers)
//is_200 = response.status_code == 200
//if is_200:
//# Get session token
//self.session_token = response.cookies.get(
//"__Secure-next-auth.session-token",
//)
//if 'json' in response.headers['Content-Type']:
//json_response = response.json()
//access_token = json_response['accessToken']
//self.save_access_token(access_token=access_token)
//self.debugger.log("SUCCESS")
//return True
//else:
//self.debugger.log(
//"Please try again with a proxy (or use a new proxy if you are using one)")
//else:
//self.debugger.log(
//"Please try again with a proxy (or use a new proxy if you are using one)")
//self.session_token = None
//self.debugger.log("Failed to get session token")
//raise Exception("Failed to get session token")