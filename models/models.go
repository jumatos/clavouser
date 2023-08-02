package models

type SecretRDSJSon struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type SignUp struct {
	USerEmail string `json:"UserEmail"`
	UserUUID  string `json:"USerUUID"`
}
