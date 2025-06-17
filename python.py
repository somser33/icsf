#!/usr/bin/env python3

import os
import time
import socket
import requests
import platform

def banner():
    os.system("clear")
    print("""
\033[1;32m
██████╗  █████╗ ███╗   ██╗ ██████╗ ██╗      █████╗ 
██╔══██╗██╔══██╗████╗  ██║██╔═══██╗██║     ██╔══██╗
██████╔╝███████║██╔██╗ ██║██║   ██║██║     ███████║
██╔═══╝ ██╔══██║██║╚██╗██║██║   ██║██║     ██╔══██║
██║     ██║  ██║██║ ╚████║╚██████╔╝███████╗██║  ██║
╚═╝     ╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝
\033[0m
    """)
    print("💻 Tool: Bangla Info Tool")
    print("👨‍💻 Coded by: SOMSER\n")

def menu():
    print("""
[1] তোমার IP দেখাও
[2] সিস্টেম ইনফো
[3] সময় দেখাও
[4] প্রস্থান করো
""")

def get_ip():
    try:
        ip = requests.get('https://api.ipify.org').text
        print(f"\n🌐 তোমার Public IP: {ip}")
    except:
        print("\n❌ ইন্টারনেট কানেকশন নেই বা সমস্যা হচ্ছে!")

def system_info():
    print("\n🛠️ সিস্টেম ইনফো:")
    print("OS:", platform.system())
    print("Version:", platform.version())
    print("Machine:", platform.machine())
    print("Processor:", platform.processor())

def show_time():
    print("\n⏰ এখন সময়:", time.ctime())

def main():
    banner()
    name = input("তোমার নাম কী? ")

    while True:
        menu()
        choice = input(f"\n{name}, তুমি কী করতে চাও? > ")

        if choice == '1':
            get_ip()
        elif choice == '2':
            system_info()
        elif choice == '3':
            show_time()
        elif choice == '4':
            print("\n👋 বিদায়, ধন্যবাদ টুলটি ব্যবহারের জন্য!")
            break
        else:
            print("\n⚠️ ভুল অপশন!")

        input("\n🔁 চালিয়ে যেতে Enter চাপো...")
        banner()

if __name__ == "__main__":
    main()
