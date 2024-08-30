#-----------------------------------------------------------------------------
# Meshplay Playground Production Environment
#-----------------------------------------------------------------------------

DEBUG=true
PROD_PLAYGROUND_PORT=80
REMOTE_PROVIDER=https://meshplay.layer5.io

#-----------------------------------------------------------------------------
# Meshplay Playground Production Deployment
#-----------------------------------------------------------------------------
.PHONY: play-prod-run play-prod-stop 

## Deploys Meshplay Playground in Prod
play-prod-run:
	(docker rm -f meshplay) || true
	docker run -d --name meshplay \
	-e PROVIDER_BASE_URLS=$(REMOTE_PROVIDER) \
	-e DEBUG=$(DEBUG) \
	-e PLAYGROUND=true \
	-e MESHPLAY_SERVER_CALLBACK_URL=https://playground.meshplay.khulnasoft.com/api/user/token \ 
	-p $(PROD_PLAYGROUND_PORT):8080 \
	layer5/meshplay:playground-latest ./meshplay
	
# docker run -d --name meshplay \
# 	-e PROVIDER_BASE_URLS=https://meshplay.layer5.io \
# 	-e PLAYGROUND=true \
# 	-p 80:8080 \
# 	layer5/meshplay:playground-latest ./meshplay

play-prod-stop:
	docker rm -f meshplay