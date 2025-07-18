#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
调用本地 hotdog 服务
python client.py zh2en "你好，世界"
python client.py en2zh "Hello world"
python client.py summary "这里是一段很长的中文文章……" --lang zh --max_len 50
"""
import argparse
import json
import sys
import requests

BASE_URL = "http://localhost:8080"

def call_api(endpoint, payload):
    url = f"{BASE_URL}/{endpoint}"
    headers = {"Content-Type": "application/json; charset=utf-8"}
    try:
        resp = requests.post(url, data=json.dumps(payload, ensure_ascii=False).encode("utf-8"), headers=headers, timeout=10)
        resp.raise_for_status()
        return resp.json().get("result", "")
    except requests.RequestException as e:
        print(f"[ERROR] {e}")
        sys.exit(1)

def main():
    parser = argparse.ArgumentParser(description="hotdog client")
    sub = parser.add_subparsers(dest="cmd", required=True)

    # 中译英
    zh2en = sub.add_parser("zh2en")
    zh2en.add_argument("text")

    # 英译中
    en2zh = sub.add_parser("en2zh")
    en2zh.add_argument("text")

    # 总结
    summary = sub.add_parser("summary")
    summary.add_argument("text")
    summary.add_argument("--lang", default="zh", help="zh 或 en")
    summary.add_argument("--max_len", type=int, default=100)

    args = parser.parse_args()

    if args.cmd == "zh2en":
        print(call_api("translate/zh2en", {"text": args.text}))
    elif args.cmd == "en2zh":
        print(call_api("translate/en2zh", {"text": args.text}))
    elif args.cmd == "summary":
        print(call_api("summarize", {"text": args.text, "lang": args.lang, "max_len": args.max_len}))

if __name__ == "__main__":
    main()