package client

import (
	pb "github.com/uz2020/petty/pb/games/xq"
	"google.golang.org/protobuf/proto"
)

func statusStream(cli *Client, argv []string) {
	stream, err := cli.gc.MyStatus(cli.ctx, &pb.MyStatusRequest{})
	if err != nil {
		pf("status stream failed %v", err)
		return
	}
	cli.ss = stream
	pf("status stream established")

	go func() {
		var msg proto.Message
		for {
			err := stream.RecvMsg(&msg)
			if err != nil {
				pl("got err ", err)
				break
			}

			pf("msg %v\n", msg)
		}
	}()
}
