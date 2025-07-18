# hotdog
零依赖 Go 后端，对接国内开源大模型（通义千问 / 智谱 GLM），提供
- 中译英
- 英译中
- 文本总结

## 快速开始
1. 克隆
   ```bash
   git clone https://github.com/bugfree5635/hotdog.git
   cd hotdog
   ```
2. 在`.env.example`编辑api-key
   ```
   # 选一种即可
   # DASHSCOPE_API_KEY=xxxxxxxxxxxxxxxx
   # ZHIPU_API_KEY=xxxxxxxxxxxxxxxx
   # OPENAI_API_KEY=xxxxxxxxxxxxxxxx
   ```
3. 复制一份`.env`
   ```bash
   cp .env.example .env
   ```   
4. `docker`构建运行服务
   ```bash
   docker build -t hotdog .
   docker run -p 8080:8080 --env-file .env hotdog
   ```
5. 客户端访问
   ```bash
   $ python client.py zh2en "你好，世界"
   Hello, World!

   $ python client.py en2zh "Hello, World!"
   你好，世界！

   $ python client.py summary "Go 是一门静态编译型语言，具有垃圾回收、并发支持等特性。" --max_len 30
   Go语言是静态编译型，支持垃圾回收和并发。
   ```
