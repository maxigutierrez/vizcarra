<?php
// Check if the form is submitted
    $req_ip = '192.168.0.10';
    $req_port = '9100';
    $req_zpl = "^XA
    ^RS8,,,3,N^FS
    ^XZ
    ^XA
    ^FO20,15^GFA,3060,3060,45,,:::::::::01NF8,... ^FS
    ^FO122,105^BX,17,200,,,,^FDUserData^FS
    ^FO85,105^A0R28,28^FDMore Data^FS
    ^FO55,105^A0R22,28^FDSome More Data^FS
    ^FO5,105^A0R50,50^FDRFID EPC^FS
    ^RR0
    ^RFW,A,0,1,3,^FDhola^FS
    ^RFR,H,0,24,2^FN1^FS^HV1,24,TAGID:,**^FS
    ^RMY^FS
    ^PQ1,0,1,Y
    ^XZ";
error_reporting(E_ALL);
/* Get the port for the service. */
$port = $req_port;
/* Get the IP address for the target host. */
$host = $req_ip;
/* construct the label */
$label = $req_zpl;
$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
if ($socket === false) {
    echo "socket_create() failed: reason: " . socket_strerror(socket_last_error    ()) . "\n";
} else {
    echo "OK.\n";
}
echo "Attempting to connect to '$host' on port '$port'...";
$result = socket_connect($socket, $host, $port);
if ($result === false) {
    echo "socket_connect() failed.\nReason: ($result) " . socket_strerror    (socket_last_error($socket)) . "\n";
} else {
    echo "OK.\n";
}
socket_set_option($socket,SOL_SOCKET, SO_RCVTIMEO, array("sec"=>5, "usec"=>0));
$result=socket_write($socket, $label, strlen($label));
echo "\nResult WRITE: ".$result;
echo "\nREAD:\n".socket_read($socket,1500)."\n";
socket_close($socket);
?>