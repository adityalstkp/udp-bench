use std::{error::Error, io, net::SocketAddr};
use tokio::net::UdpSocket;

struct UdpServer {
    socket: UdpSocket,
    buff: Vec<u8>,
    print_out: Option<(usize, SocketAddr)>,
}


impl UdpServer {
    async fn start(self) -> Result<(), io::Error> {
        let UdpServer {
            socket,
            mut buff,
            mut print_out,
        } = self;

        loop {
            if let Some((size, peer)) = print_out {
                let _amt = socket.send_to(&buff[..size], &peer).await?;
            }

            print_out = Some(socket.recv_from(&mut buff).await?);
        }
    }

}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let addr = "127.0.0.1:3001";
    let socket = UdpSocket::bind(&addr).await?;

    println!("Listening on {}", socket.local_addr()?);

    let udp_server = UdpServer{
        socket,
        buff: vec![0; 1024],
        print_out: None,
    };

    udp_server.start().await?;
    Ok(())
}
