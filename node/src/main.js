const udp = require('dgram')

/**
    @return {import('dgram').Socket}
*/
function createUdpServer(datagram) {
    const udpServer = udp.createSocket('udp4')
    udpServer.bind(3002, '0.0.0.0', function() {
        udpServer.setRecvBufferSize(datagram)
    })

    return udpServer 
}

// implement worker_threads??
function main() {
    const udpS = createUdpServer(1024)

    udpS.on('listening', function () {
        const addr = udpS.address()
        console.log(`Listening on ${addr.port}`) 
    })

    udpS.on('message', function(msg) {
        console.log(msg.toString())
    })
}

main()


