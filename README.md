# ootbat

I keep tunnel visioning on code and then my battery dies and I get pissed off. This program plays the low heart sound from Ocarina of Time every ten seconds if my battery falls below 5%. This way I can feel nostalgic while not get pissed off anymore. If your laptop has an intel processor, one internal battery, and runs Linux this program might also work for you.

![hearts](hearts.png)

```
$ sudo bash install.sh
$ systemctl status ootbat
● ootbat.service - ootbat
     Loaded: loaded (/usr/lib/systemd/system/ootbat.service; enabled; vendor preset: disabled)
     Active: active (running) since Tue 2020-10-27 15:23:53 CDT; 8min ago
   Main PID: 21212 (ootbat)
      Tasks: 5 (limit: 38061)
     Memory: 1.1M
     CGroup: /system.slice/ootbat.service
             └─21212 /opt/ootbat/ootbat

Oct 27 15:23:53 xps systemd[1]: Started ootbat.
```