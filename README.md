# Announcer System
Repository for Announcer System to help Cloud Native Indonesia community announce something (meetup, kubeweekly, etc.) directly into community member.

```
make deps
```

## Kubeweekly 
```
export USERNAME=""
export PASSWORD=""

make kubeweekly-build
make kubeweekly-run
```
## CNCF Newsroom
```
export USERNAME=""
export PASSWORD=""

make cncf-newsroom-build
make cncf-newsroom-run
```
## CNCF Webinar
```
export USERNAME=""
export PASSWORD=""

make cncf-webinar-build
make cncf-webinar-run
```
## Orchestrator
```
export USERNAME=""
export PASSWORD=""
export TELEGRAM_TOKEN=""
export TELEGRAM_CHATID=""

make orchestrator-build
cd cmd/orchestrator
./orchestrator
```

