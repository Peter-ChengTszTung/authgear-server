# export K6_HTTP_DEBUG=true
# We cannot use K6_DURATION and K6_VUS because
# using them will make k6 to ignore scenarios.
export ENDPOINT=http://accounts.portal.localhost:3100
export CLIENT_ID=9fba9e479d3a3863
# export REDIRECT_URI=http://localhost:8000/oauth-redirect
export REDIRECT_URI=com.example://host
export FIXTURE_FIXED_OTP=000000
export FIXTURE_PASSWORD=a12345678
export FIXTURE_EMAIL_DOMAIN=example.com
export FIXTURE_PHONE_NUMBER_COUNTRY_CALLING_CODE=+852
export FIXTURE_PHONE_NUMBER_LOW=52006000
export FIXTURE_PHONE_NUMBER_HIGH=57999999

.PHONY: go-test
go-test:
	go test .

.PHONY: k6
k6:
	xk6 build --with github.com/authgear/authgear-server/k6=.

.PHONY: clean
clean:
	rm -rf ./k6

.PHONY: signup
signup:
	./k6 run --config options.json ./signup.js

.PHONY: login
login:
	./k6 run --config options.json ./login.js

.PHONY: refresh_access_token
refresh_access_token:
	./k6 run --config options.json ./refresh_access_token.js

.PHONY: get_user_info
get_user_info:
	./k6 run --config options.json ./get_user_info.js

.PHONY: biometric_login
biometric_login:
	./k6 run --config options.json ./biometric_login.js

.PHONY: healthz
healthz:
	./k6 run --config options.json ./healthz.js
