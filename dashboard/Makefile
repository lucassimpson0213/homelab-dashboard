build: 
	./backend/scripts/build-ts-and-go.sh

test:
	go test ./...

lint:
	go vet ./...

dev:
	systemctl --user restart homelab-dashboard-dev.service

logs:
	journalctl --user -u homelab-dashboard-dev.service -f

restart:
	systemctl --user restart homelab-dashboard-dev.service





