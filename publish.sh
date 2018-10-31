# Publish on surge.sh using zed@nsdmdh.com
# Recc requirements:
# - Node 8
# - surge
# - npx
hugo &&\
    echo "zeddee.com" > docs/CNAME &&\
    npx surge docs zeddee.com

