build:
        go build main.go

deploy:
        sudo systemctl stop gs_tbot.service
        sudo cp ./gs_tbot.service /lib/systemd/system/gs_tbot.service
        sudo systemctl daemon-reload
        sudo systemctl start gs_tbot.service