# tcp三次握手

```
tcpdump -S tcp -i eth5 -nn port 80 and host 19.19.19.70 and host 19.19.19.104


22:12:36.232179 IP 19.19.19.104.59327 > 19.19.19.70.80: Flags [S], seq 1817773983, win 64240, options [mss 1460,nop,wscale 8,nop,nop,sackOK], length 0
22:12:36.232227 IP 19.19.19.70.80 > 19.19.19.104.59327: Flags [S.], seq 1737731967, ack 1817773984, win 64240, options [mss 1460,nop,nop,sackOK,nop,wscale 7], length 0
22:12:36.232373 IP 19.19.19.104.59327 > 19.19.19.70.80: Flags [.], ack 1737731968, win 8212, length 0

22:12:36.243523 IP 19.19.19.104.59327 > 19.19.19.70.80: Flags [P.], seq 1817773984:1817774066, ack 1737731968, win 8212, length 82: HTTP: GET /api.php HTTP/1.1
22:12:36.243584 IP 19.19.19.70.80 > 19.19.19.104.59327: Flags [.], ack 1817774066, win 502, length 0
22:12:36.248572 IP 19.19.19.70.80 > 19.19.19.104.59327: Flags [P.], seq 1737731968:1737732418, ack 1817774066, win 502, length 450: HTTP: HTTP/1.1 200 OK

22:12:36.248974 IP 19.19.19.104.59327 > 19.19.19.70.80: Flags [F.], seq 1817774066, ack 1737732418, win 8210, length 0
22:12:36.249070 IP 19.19.19.70.80 > 19.19.19.104.59327: Flags [F.], seq 1737732418, ack 1817774067, win 502, length 0
22:12:36.249339 IP 19.19.19.104.59327 > 19.19.19.70.80: Flags [.], ack 1737732419, win 8210, length 0

tcp:
Client > Server (seq = x,syn = 1,stauts=SYN_SEND)
Server > Client (ack =x+1,seq = y,syn = 1,stauts=SYN_RECV)
Client > Server (ack = y+1,Client和Server同步进入status=ESTABLISHED)

http:
client > Server (Method = get)
Server > Client (ack确认)
Server > Client (数据包和状态码)
Client > Server (ack确认)

tcp:
Client > Server (发送fin包,seq = n,ack = i)
Server > Client (发送fin包,ack = n+1)
Server > Client (发送fin包,seq = i,ack = n+1)
Client > Server (ack=i+1)
```