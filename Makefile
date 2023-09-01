build:
	go build -o bin/socket_server

edit_conntrack:
	./sh_utils/edit_conntrack_table

install_golang:
	chmod +xf ./sh_utils/install_golang.sh
	./sh_utils/install_golang.sh