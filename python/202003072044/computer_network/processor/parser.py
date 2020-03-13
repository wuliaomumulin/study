# -*- encoding=utf-8

import struct
import socket

# IP报文解析器
class IPParser:
	IP_HEADER_LENGTH =20

	@classmethod
	def parse_ip_header(cls,ip_header):
		'''
		IP 报文格式
		1. 4位ip_version 4位IP头部长度 8位服务类型 16位总长度
		2. 16位标识符 3位标记位 3位片偏移 
		3. 8位TTL 8位协议 16位IP头校验和
		4. 32位源IP地址
		5. 32位目的IP地址
		: param ip_header
		: return:
		'''
		# 第一行解析，前四个字节交给它解析
		line1 = struct.unpack('>BBH',ip_header[:4])
		'''
		取出4个byte流，由于1Byte等于8bit=8位，而一个ip_version仅仅只有4位bit,
		我们需要取出第一个byte的一半
		所有我们需要将数据右移四位
		如
		eg:line[0] >> 4 == 11110000 -> 00001111
		eg:11111111 & 00001111 ==  00001111
		'''
		ip_version = line1[0] >> 4
		iph_length = line1[0] & 15 * 4 # 为了方便我们将IP头部乘以4
		pkg_length = line1[2] #直接获取总长度

		# 第三行解析
		line3 = struct.unpack('>BBH',ip_header[8:12])
		TTL = line3[0]
		protocol = line3[1]
		iph_checksum = line3[2]
		# 第四行解析
		line4 = struct.unpack('>4s',ip_header[12:16])
		# 将字节流解析成IP
		src_ip = socket.inet_ntoa(line4[0])
		# 第五行解析
		line5 = struct.unpack('>4s',ip_header[16:20])
		dst_ip = socket.inet_ntoa(line5[0])

		return {
		  'ip_version':ip_version,
		  'iph_length':iph_length,
		  'pkg_length':pkg_length,
		  'TTL':TTL,
		  'protocol':protocol,#协议
		  'iph_checksum':iph_checksum,#校验和
		  'src_ip':src_ip,
		  'dst_ip':dst_ip
		}

	@classmethod
	def parse(self,packet):
		# 前20字节
		ip_header = packet[:20]
		return self.parse_ip_header(ip_header)
