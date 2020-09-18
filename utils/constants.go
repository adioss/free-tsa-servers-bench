package utils

const JarName = "generated.jar"

var TsaServers = [...]string{
	"http://timestamp.digicert.com",
	"http://timestamp.comodoca.com/rfc3161",
	"http://timestamp.entrust.net/TSS/RFC3161sha2TS",
	"http://sha256timestamp.ws.symantec.com/sha256/timestamp",
	"http://timestamp.globalsign.com/scripts/timstamp.dll",
	"http://time.certum.pl",
	"http://rfc3161timestamp.globalsign.com/advanced",
	//"http://zeitstempel.dfn.de", // LTV ko
	//"http://timestamp.apple.com/ts01", // LTV ko
	//"https://freetsa.org", // LTV ko
	//"http://tsa.wotrus.com",           // unavailable after 100 hints
	//"http://tsa.startssl.com/rfc3161", // unavailable after 100 hints
	//"https://timestamp.geotrust.com/tsa",
	//"http://timestamp.wosign.com",
	//"https://ca.signfiles.com/tsa/get.aspx",
	//"http://services.globaltrustfinder.com/adss/tsa",
	//"https://tsp.iaik.tugraz.at/tsp/TspRequest",
	//"http://dse200.ncipher.com/TSS/HttpTspServer",
	//"http://tsa.safecreative.org",
}
