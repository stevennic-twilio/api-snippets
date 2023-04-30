using System;
using Twilio.TwiML;
using Twilio.TwiML.Voice;


class Example
{
    static void Main()
    {
        var response = new VoiceResponse();
        var dial = new Dial();
        dial.Conference("NoMusicNoBeepRoom", beep: Conference.BeepEnum.False,
            waitUrl: new Uri("http://your-webhook-host.com"),
            startConferenceOnEnter: true, endConferenceOnExit: true);
        response.Append(dial);

        Console.WriteLine(response.ToString());
    }
}