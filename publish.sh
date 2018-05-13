# Publish on surge.sh using zed@nsdmdh.com
# Recc requirements:
# - Node 8
# - surge
# - npx
hugo &&\
    echo "zeddee.com" > public/CNAME &&\
    npx surge public zeddee.com

