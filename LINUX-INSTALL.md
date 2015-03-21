# Installing Ubuntu Server with Chrome Browser How-To

## Create Bootable USB Drive from your Mac

Download Ubuntu Server for amd64 http://www.ubuntu.com/download/server (on a Mac).

Plug in a USB flash drive.

Convert the .iso file to .img using the convert option of hdiutil e.g.,

<code>
hdiutil convert -format UDRW -o ~/path/to/target.img ~/path/to/ubuntu.iso
</code>

Run

<code>
diskutil list
</code>

again and determine the device node assigned to your flash media (e.g. `/dev/disk2`).

Run

<code>
diskutil unmountDisk /dev/diskN
</code>

(replace N with the disk number from the last command; in the previous example, N would be 2).

Execute

<code>
sudo dd if=/path/to/downloaded.img of=/dev/rdiskN bs=1m
</code>

There should be a Mac OS X Finder prompt, at which point you can eject.


References: http://www.ubuntu.com/download/desktop/create-a-usb-stick-on-mac-osx

## Boot up Family Room Pro from the USB Stick

Plug in the USB drive, and right after hitting the start button, hit the `del` key to load up the firmware.

Select the USB key as the first two boot options.

Save & Exit firmware, to boot.

## Run Ubuntu Installer

Install linux, and select entire hard drive `with LVM`.  Select `100%` of your drive to give it to use.  

Select SSH server to install, which makes saving odd video driver issues easier to get to.

Let install finish, and reboot

## Fix `swap` issue (actually a kernel video issue)

Boot into recovery mode (you should have a 2 second chance to select this in grub2), and edit grub settings:

<code>
nano /etc/default/grub
</code>

(`ctrl-o` and then `ctrl-x` to save and quit)

Add `nomodeset` to `GRUB_CMDLINE_LINUX_DEFAULT=""` and `GRUB_CMDLINE_LINUX=""` lines, e.g.

<code>
GRUB_CMDLINE_LINUX_DEFAULT="nomodeset"
GRUB_CMDLINE_LINUX="nomodeset"
</code>

Then run `update-grub` to tell grub to use these options.

<code>
update-grub
</code>

References: 
http://blog.jamesrhall.com/2014/05/update-grub-2-options.html
http://blog.jamesrhall.com/2014/04/ubuntu-server-1404-fun.html
http://serverfault.com/questions/546079/ubuntu-server-hanging-on-adding-swap



## Fix `apt-get` hanging issue

Edit GAI config (?): open `/etc/gai.conf`

<code>
nano /etc/gai.conf
</code>

change line ~54 to *uncomment* out the following:

<code>
precedence ::ffff:0:0/96  100
</code>

(`ctrl-o` and then `ctrl-x` to save and quit)

References: http://askubuntu.com/questions/574569/apt-get-stuck-at-0-connecting-to-us-archive-ubuntu-com

## Reboot

reboot or sudo reboot.

## Update OS

Do the following:

<code>
sudo apt-get update
sudo apt-get upgrade
sudo apt-get dist-upgrade
</code>

## Install NVIDIA Drivers

<code>
sudo apt-get install nvidia-331
</code>

## Install Desktop (xfce4)

<code>
sudo apt-get install xubuntu-desktop
</code>


## Install Chrome

From a USB Stick?  From the web?

