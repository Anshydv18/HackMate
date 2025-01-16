package templates

import (
	"context"
	"fmt"
)

func GetCustomisedMessage(ctx *context.Context, to string, by string, details string) string {
	return fmt.Sprintf(`Dear %s,

We have received a request from %s who would like to connect with you. Here are their details for your reference:

%s

If you are open to this connection, please feel free to reach out to them directly.

Should you have any questions or concerns, do not hesitate to contact us.

Best regards,  
HackMate`, to, by, details)
}

func OtpVerificationTemplate(ctx *context.Context, otp int64) string {
	return fmt.Sprintf(`Dear User,
	
Please Verify Your account 
Otp  %d

Best regards,  
HackMate	`, otp)
}
