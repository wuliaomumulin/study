# -*- encoding=utf-8

import struct
import socket

class TransParser:
	# IP头部偏移20bit
	IP_HEADER_OFFSET = 20
	# UDP头部长度
	UDP_HEADER_LENGTH = 8
	# TCP头部长度
	TCP_HEADER_LENGTH = 20

	# 报文里面的数据
	def data2str(data):
		l = len(data)
		data = struct.unpack(l*'B',data)
		string = ''
		for ch in data:
			if ch>127 or ch<32:
				string += '.'
			else:
				string += chr(ch)
		return string

# UDP报文解析器
class UDPParser(TransParser):

	'''
	1.16位源端口 16位目标端口
	2、16UDP长度 16位校验和
	3、32位UDP数据
	'''
	@classmethod
	def parse_udp_header(cls,udp_header,packet):

		udp_header = struct.unpack('>HHHH',udp_header)

		return {
		  'src_port':udp_header[0],
		  'dst_port':udp_header[1],
		  'udp_length':udp_header[2],
		  'udph_checknum':udp_header[3],
		  'data' : super().data2str(packet[cls.IP_HEADER_OFFSET+cls.UDP_HEADER_LENGTH:])
		}

	@classmethod
	def parse(self,packet):
		# 前20字节
		udp_header = packet[self.IP_HEADER_OFFSET:self.IP_HEADER_OFFSET+self.UDP_HEADER_LENGTH]
		return self.parse_udp_header(udp_header,packet)


# TCP报文解析器
class TCPParser(TransParser):
	
	'''
	TCP header 结构
	1.16位源端口 16位目的端口
	2.序列号
	3.确认号
	4.4位数据偏移 6位保留字段 6位标志位 16位窗口大小
	5.16位校验和 16位紧急指针

	单位换算:
	一个字节=1byte=8位,四个字节=32位
	一个字=2byte=16位
	'''
	@classmethod
	def parse_tcp_header(cls,tcp_header,packet):

		header = {
			"flag":{}
		}

		line1 = struct.unpack('>HH',tcp_header[:4])
		header['src_port'] = line1[0]
		header['dst_port'] = line1[1]
		line2 = struct.unpack('>L',tcp_header[4:8])
		header['seq_num'] = line2[0]
		line3 = struct.unpack('>L',tcp_header[8:12])
		header['ack_num'] = line3[0]
		line4 = struct.unpack('>BBH',tcp_header[12:16])
		header['data_offset'] = line4[0] >> 4
		flags = line4[1] & int('00111111',2)
		header['flag']['FIN'] = flags & 1
		header['flag']['SYN'] = (flags >> 1) & 1
		header['flag']['RST'] = (flags >> 2) & 1
		header['flag']['PSH'] = (flags >> 3) & 1
		header['flag']['ACK'] = (flags >> 4) & 1
		header['flag']['URG'] = (flags >> 5) & 1
		header['win_size'] = line4[2]
		line5 = struct.unpack('>HH',tcp_header[16:20])
		header['check_num'] = line5[0]
		header['urg_point'] = line5[1]
		# 数据体
		header['data'] = super().data2str(packet[cls.IP_HEADER_OFFSET+cls.TCP_HEADER_LENGTH:])
		return header
		
	@classmethod
	def parse(self,packet):
		# 前20字节
		tcp_header = packet[self.IP_HEADER_OFFSET:self.IP_HEADER_OFFSET+self.TCP_HEADER_LENGTH]
		return self.parse_tcp_header(tcp_header,packet)
