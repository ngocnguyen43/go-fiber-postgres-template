package utils

import "time"

const MaxOpenConnections = 50
const MaxWaitConnections = 1000
const MaxReadTimeOut = 30 * time.Second
const MaxWriteTimeOut = 30 * time.Second
const HealthyMessage = "It's healthy"
const PrivateKeyPath = "keys/private_key.pem"
const PublicKeyPath = "keys/public_key.pem"
