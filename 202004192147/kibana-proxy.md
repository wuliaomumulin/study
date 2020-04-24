把一下内容加到apache2.conf 使用的配置文件的 VirtualHost标签内


    Proxyrequests off
    <Proxy *>
            Order deny,allow
            Allow from all
    </Proxy>

    RewriteEngine on
    RewriteRule ^ /kibana/(.*)$ /ossim/test_kibana/kibana.php?url=$1 [QSA,P]
    ProxyPassReverse /kibana/ http://127.0.0.1:5601/kibana/

    SSLProxyEngine on
    SSLProxyVerify none
    SSLProxyCheckPeerCN off
    SSLProxyCheckPeerName off
    SSLProxyCheckPeerExpire off





# 20200424

    ProxyRequests on
    <Proxy *>
            Order deny,allow
            Allow from all
    </Proxy>

    ProxyPreserveHost On

        ProxyPass /elasticsearch/ http://192.168.1.200:5601/elasticsearch/
            ProxyPassReverse /elasticsearch/ http://192.168.1.200:5601/elasticsearch/

            ProxyPass /app/kibana http://192.168.1.200:5601/app/kibana
            ProxyPassReverse /app/kibana http://192.168.1.200:5601/app/kibana

        ProxyPass /bundles/ http://192.168.1.200:5601/bundles/
            ProxyPassReverse /bundles/ http://192.168.1.200:5601/bundles/

        ProxyPass /built_assets/ http://192.168.1.200:5601/built_assets/
            ProxyPassReverse /built_assets/ http://192.168.1.200:5601/built_assets/


        ProxyPass /ui/ http://192.168.1.200:5601/ui/
            ProxyPassReverse /ui/ http://192.168.1.200:5601/ui/


        ProxyPass /node_modules/ http://192.168.1.200:5601/node_modules/
            ProxyPassReverse /node_modules/ http://192.168.1.200:5601/node_modules/

        ProxyPass /translations/ http://192.168.1.200:5601/translations/
            ProxyPassReverse /translations/ http://192.168.1.200:5601/translations/

         ProxyPass /api/ http://192.168.1.200:5601/api/
            ProxyPassReverse /api/ http://192.168.1.200:5601/api/

        ProxyPass /internal/ http://192.168.1.200:5601/internal/
            ProxyPassReverse /internal/ http://192.168.1.200:5601/internal/


        ProxyPass /goto/ http://192.168.1.200:5601/goto/
            ProxyPassReverse /goto/ http://192.168.1.200:5601/goto/
