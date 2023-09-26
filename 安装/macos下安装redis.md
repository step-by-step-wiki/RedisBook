# macos下安装redis

通过二进制的方式安装 Redis 的步骤如下：

1. **下载 Redis 源码**:
   你可以从 Redis 官方网站上下载最新的源码包。由于此环境中无法访问外部网站，你可以自行前往 Redis 官方网站或其 GitHub 仓库进行下载。

2. **解压源码包**:
   假设你下载了 `redis-6.2.5.tar.gz` (版本号可能会不同)，在终端中进入到下载目录并执行：

   ```bash
   tar xzf redis-6.2.5.tar.gz
   cd redis-6.2.5
   ```

3. **编译 Redis**:

   在 `redis-6.2.5` 目录中，执行以下命令来编译 Redis：

   ```bash
   make
   ```

   编译完成后，`src` 目录下会生成相关的二进制文件，如 `redis-server` 和 `redis-cli`。

4. **测试编译结果**:

   在 `redis-6.2.5` 目录中，执行：

   ```bash
   make test
   ```

   这会运行 Redis 的测试套件，确保编译的版本没有问题。

5. **安装**:

   如果你想将 Redis 的二进制文件安装到系统中，可以执行：

   ```bash
   sudo make install
   ```

   默认情况下，这会将 Redis 安装到 `/usr/local/bin` 目录。

6. **启动 Redis**:

   你可以使用以下命令启动 Redis：

   ```bash
   redis-server
   ```

   如果你希望使用自定义的配置文件启动 Redis，可以这样做：

   ```bash
   redis-server /path/to/your/redis.conf
   ```

7. **测试 Redis**:

   使用 Redis 客户端工具 `redis-cli` 来测试 Redis 服务器：

   ```bash
   redis-cli ping
   ```

   如果 Redis 服务器正在运行，你应该会看到 `PONG` 作为响应。

以上就是通过二进制方式在 macOS 上安装 Redis 的步骤。希望这可以帮到你！