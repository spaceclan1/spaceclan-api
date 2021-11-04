repository="zura101"
image_name="spaceclan-data-gatherer"
TIMESTAMP=`date "+%Y.%m.%d.%H%M%S"`

docker build -t $repository/$image_name:$TIMESTAMP .
echo $repository/$image_name:$TIMESTAMP