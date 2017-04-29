# payirGolang
This is a simple GoLang package for pay.ir gateway.

For requesting a payment transId use the following code sample:

<pre>
  p := &payRequest{
		APIKey:       APIKEY,
		Amount:       50000,
		Redirect:     RedirectTo,
		FactorNumber: "1369",
	}
	pr := p.RequestPaymentToken()
	fmt.Println(pr)
</pre>
  
  
  And here is how you can verify the transaction :
  
  <pre>
  v := &verifyRequest{
		APIKEY:  APIKEY,
		TransID: YOURTRANSID,
	}
	vr := v.RequestPaymentVerification()
	fmt.Println(vr)
  </pre>
  
  Any suggestion or improvement regarding the way this package is coded will be apperciated.
  This is my first golang project so it's not perfect.
