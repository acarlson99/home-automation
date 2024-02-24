set -ex

useradd go-home

sudo mkdir /home/go-home
sudo chown -R go-home /home/go-home/

mkdir -p /tmp/go-home-install
cd /tmp/go-home-install

cat <<EOF >>/tmp/go-home-install/go-home.service
[Unit]
Description=Home Automation Service
After=network.target

[Service]
Type=simple
ExecStart=/home/go-home/go/bin/home-automation -devices=/etc/go-home/devices.textpb -schedule=/etc/go-home/schedule.textpb -log-level=info
Restart=on-failure
User=go-home

ProtectSystem=full
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

sudo mv /tmp/go-home.service /lib/systemd/system/go-home.service

if ! command -v go &>/dev/null; then
    # NOTE: this depends on cpu arch
    wget https://go.dev/dl/go1.22.0.linux-arm64.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.0.linux-arm64.tar.gz

    cat <<EOF >>/home/go-home/.bashrc
export GOROOT=/usr/local/go
export GOPATH=/home/go-home/go
export PATH=\$PATH:\$GOPATH/bin:\$GOROOT/bin
EOF
    source /home/go-home/.bashrc

    # raspberry-pi fix
    go env -w GOARCH=arm
fi

go install github.com/acarlson99/home-automation@latest

chmod -R +w /home/go-home/go/

pushd /home/go-home/go/pkg/mod/github.com/acarlson99/home-automation@*
go test ./...
popd

# places binary in /home/go-home/go/bin/
make -C /home/go-home/go/pkg/mod/github.com/acarlson99/home-automation@* install
rm -rf /home/go-home/go/

# configs live in /etc/go-home/
mkdir /etc/go-home
cp *.textpb /etc/go-home/
chown -R go-home /etc/go-home/

cp /tmp/go-home-install/go-home.service /lib/systemd/system/go-home.service
systemctl daemon-reload
systemctl enable go-home
systemctl start go-home
