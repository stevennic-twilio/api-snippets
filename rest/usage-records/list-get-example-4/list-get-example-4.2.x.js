// Download the Node helper library from twilio.com/docs/node/install
// These vars are your accountSid and authToken from twilio.com/user/account
var accountSid = 'ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX';
var authToken = "your_auth_token";
var client = require('twilio')(accountSid, authToken);

client.usage.records.daily.list({ category: "calls-inbound",
    startDate: "2012-09-01",
    endDate: "2012-09-30"
  }, function(err, data) {
    data.usageRecords.forEach(function(record) {
        console.log(record.price);
    });
});