build:
	rm -rf ./libvirt-go-examples
	go build .

run:
	rm -rf ./images/vm-images/vm2/user-data.img
	genisoimage -output ./images/vm-images/vm2/user-data.img -volid cidata -joliet -rock ./images/user-data/vm2/user-data ./images/user-data/vm2/meta-data
	./libvirt-go-examples
