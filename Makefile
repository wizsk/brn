PACKAGE_NAME = brn
VERSION = 1
ARCHITECTURE = amd64
EXEC_DIR = build/
MAINTAINER = https://github.com/wizsk
DESCRIPTION =  Bulk rename files in a sweep with your favourite text editor

build:
	@echo
	@echo "[+] Version: $(VERSION)"
	@echo

	@mkdir -p $(EXEC_DIR)
	# @mkdir -p $(DEB_PACKAGE_DIR)
	# @mkdir -p $(DEBIAN_DIR)
	# @mkdir -p $(TARGET_EXECUTABLE_PATH)

	@echo "[+] Building the Linux version"
	@go build -ldflags "-s -w" -o $(EXEC_DIR)brn

	@echo "[+] Packaging the Linux version"
	@tar -czvf $(EXEC_DIR)brn_Linux.tar.gz -C $(EXEC_DIR) brn > /dev/null
	@sha256sum $(EXEC_DIR)brn_Linux.tar.gz
	# @sha256sum $(EXEC_DIR)brn_Linux.tar.gz > $(EXEC_DIR)brn_Linux_sha256sum.txt

	@echo
	@echo "[+] Building the static Linux version"
	@env GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o $(EXEC_DIR)brn

	@echo "[+] Packaging the static Linux version"
	@tar -czvf $(EXEC_DIR)brn_Linux_static.tar.gz -C $(EXEC_DIR) brn > /dev/null
	@sha256sum $(EXEC_DIR)brn_Linux_static.tar.gz
	# @sha256sum $(EXEC_DIR)brn_Linux_static.tar.gz > $(EXEC_DIR)brn_Linux_static_sha256sum.txt


	@echo "[+] Removing the static Linux binary"
	@rm $(EXEC_DIR)brn


	@echo
	@echo "[+] Building the Windows version"
	@env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(EXEC_DIR)brn.exe

	@echo "[+] Packaging the Windows version"
	@zip -j $(EXEC_DIR)brn_Windows.zip $(EXEC_DIR)brn.exe > /dev/null
	@sha256sum  $(EXEC_DIR)brn_Windows.zip
	# @sha256sum  $(EXEC_DIR)brn_Windows.zip > $(EXEC_DIR)brn_Windows_sha256sum.txt

	@echo "[+] Removing the Windows binary"
	@rm $(EXEC_DIR)brn.exe

	@echo
	@echo "[+] Building the MacOS version"
	@env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $(EXEC_DIR)brn

	@echo "[+] Packaging the MacOS version"
	@tar -czvf $(EXEC_DIR)brn_MacOS.tar.gz -C $(EXEC_DIR) brn > /dev/null
	@sha256sum $(EXEC_DIR)brn_MacOS.tar.gz
	# @sha256sum $(EXEC_DIR)brn_MacOS.tar.gz > $(EXEC_DIR)brn_MacOS_sha256sum.txt


	@echo "[+] Removing the MacOS binary"
	@rm $(EXEC_DIR)brn

	@echo
	@echo "[+] Building the MacOS ARM version"
	@env GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o $(EXEC_DIR)brn

	@echo "[+] Packaging the MacOS ARM version"
	@tar -czvf $(EXEC_DIR)brn_MacOS_ARM.tar.gz -C $(EXEC_DIR) brn > /dev/null
	@sha256sum $(EXEC_DIR)brn_MacOS_ARM.tar.gz
	# @sha256sum $(EXEC_DIR)brn_MacOS_ARM.tar.gz > $(EXEC_DIR)brn_MacOS_ARM_sha256sum.txt

	@echo "[+] Removing the MacOS ARM binary"
	@rm $(EXEC_DIR)brn

	@echo
	@echo "[+] Building the FreeBSD version"
	@env GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -w" -o $(EXEC_DIR)brn

	@echo "[+] Packaging the FreeBSD AMD64 version"
	@tar -czvf $(EXEC_DIR)brn_FreeBSD.tar.gz -C $(EXEC_DIR) brn > /dev/null
	@sha256sum $(EXEC_DIR)brn_FreeBSD.tar.gz
	# @sha256sum $(EXEC_DIR)brn_FreeBSD.tar.gz > $(EXEC_DIR)brn_FreeBSD_sha256sum.txt

	@echo "[+] Removing the FreeBSD binary"
	@rm $(EXEC_DIR)brn

	@echo
	@echo "[+] Done"
