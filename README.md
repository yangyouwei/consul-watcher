# consul-watcher

### 简介

    本程序接收consul watch 监控到的变化（服务--json）
    解析json，使用go的模板库。修改nginx upstream配置文件。
    调用dyups的http接口，修改nginx 内存中的upstream的配置，不用重启nginx

使用环境

    consul  使用了服务注册，使用了watch功能。
        watch监控服务变化，向指定api推送可用服务（json格式）。 
    tengine（编译了dyups模块）
        开启dyups模块的http接口
        
tengine 编译配置

    ./configure --add-module=/root/tengine-2.3.2/modules/ngx_http_upstream_dyups_module/

    [root@localhost conf.d]# cat dyups_management.conf 
         server {
            listen  18882; 
            location / {
                dyups_interface;
            }
        }

### config

    [main]
    listen_port=8000
    #是否存log为文件
    log=false
    #不指定目录，就再程序当前目录下创建processname.log
    log_file_dir=./log
    
    [ups]
    #模板路径
    tpl_path=./ups.tpl
    #upstream路径，必须以“/”结尾 例如：/etc/nginx/conf.d/
    upstream_path=./
    #dyups模块的接口地址，必须以“/”结尾 例如http://127.0.0.1:18882/upstream/
    dyups_url=http://192.168.3.112:18882/upstream/
