APP = dean

deploy:
	 flyctl deploy -a ${APP}
conf secret:
	flyctl secrets import -a ${APP} < .env.${APP}
bash:
	flyctl ssh console -a  ${APP} -C /bin/bash

build:
	@bash ./build/util.sh until::build

.PHONY:  build