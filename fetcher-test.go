package main

func main() {

	var weatherCanadaVancouverFeedXML = []byte(`
<feed xmlns="http://www.w3.org/2005/Atom" xml:lang="en-ca">
    <title>Vancouver - Weather - Environment Canada</title>
    <link rel="related" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html" type="text/html"/>
    <link rel="self" href="https://www.weather.gc.ca/rss/city/bc-74_e.xml" type="application/atom+xml"/>
    <link rel="alternate" hreflang="fr-ca" href="https://www.meteo.gc.ca/rss/city/bc-74_f.xml" type="application/atom+xml"/>
    <author>
        <name>Environment Canada</name>
        <uri>https://www.weather.gc.ca</uri>
    </author>
    <updated>2018-10-09T04:11:03Z</updated>
    <id>tag:weather.gc.ca,2013-04-16:20181009041103</id>
    <logo>https://www.weather.gc.ca/template/gcweb/v4.0.24/assets/wmms-alt.png</logo>
    <icon>https://www.weather.gc.ca/template/gcweb/v4.0.24/assets/favicon.ico</icon>
    <rights>Copyright 2018, Environment Canada</rights>
    <entry>
        <title>No watches or warnings in effect, Vancouver</title>
        <link type="text/html" href="https://www.weather.gc.ca/warnings/index_e.html"/>
        <updated>2018-10-01T07:00:00Z</updated>
        <published>2018-10-01T07:00:00Z</published>
        <category term="Warnings and Watches"/>
        <summary type="html">No watches or warnings in effect.</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_w1:20181001070000</id>
    </entry>
    <entry>
        <title>Current Conditions: Mostly Cloudy, 10.5&#xB0;C</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-09T04:08:00Z</updated>
        <published>2018-10-09T04:08:00Z</published>
        <category term="Current Conditions"/>
        <summary type="html"><![CDATA[<b>Observed at:</b> Vancouver Int'l Airport 9:08 PM PDT Monday 08 October 2018 <br/>
<b>Condition:</b> Mostly Cloudy <br/>
<b>Temperature:</b> 10.5&deg;C <br/>
<b>Pressure / Tendency:</b> 101.2 kPa falling<br/>
<b>Visibility:</b> 12.9 km<br/>
<b>Humidity:</b> 97 %<br/>
<b>Dewpoint:</b> 10.0&deg;C <br/>
<b>Wind:</b> WNW 7 km/h<br/>
<b>Air Quality Health Index:</b> 2 <br/>
]]>        </summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_cc:20181009040800</id>
    </entry>
    <entry>
        <title>Monday night: Periods of drizzle. Temperature steady near 11. POP 60%</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Cloudy. 60 percent chance of drizzle this evening. Temperature steady near 11. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc1:20181008230000</id>
    </entry>
    <entry>
        <title>Tuesday: Clearing. High 15.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Clearing in the morning. Fog patches dissipating in the morning. High 15. UV index 3 or moderate. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc2:20181008230000</id>
    </entry>
    <entry>
        <title>Tuesday night: A few clouds. Low 8.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Clear. Becoming partly cloudy after midnight. Fog patches developing after midnight. Low 8. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc3:20181008230000</id>
    </entry>
    <entry>
        <title>Wednesday: Sunny. High 15.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Sunny. High 15. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc4:20181008230000</id>
    </entry>
    <entry>
        <title>Wednesday night: Clear. Low 7.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Clear. Low 7. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc5:20181008230000</id>
    </entry>
    <entry>
        <title>Thursday: Sunny. High 13.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Sunny. High 13. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc6:20181008230000</id>
    </entry>
    <entry>
        <title>Thursday night: Cloudy periods. Low 9.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Cloudy periods. Low 9. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc7:20181008230000</id>
    </entry>
    <entry>
        <title>Friday: Sunny. High 15.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Sunny. High 15. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc8:20181008230000</id>
    </entry>
    <entry>
        <title>Friday night: Clear. Low 6.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Clear. Low 6. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc9:20181008230000</id>
    </entry>
    <entry>
        <title>Saturday: Sunny. High 15.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Sunny. High 15. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc10:20181008230000</id>
    </entry>
    <entry>
        <title>Saturday night: Clear. Low plus 5.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Clear. Low plus 5. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc11:20181008230000</id>
    </entry>
    <entry>
        <title>Sunday: Sunny. High 15.</title>
        <link type="text/html" href="https://www.weather.gc.ca/city/pages/bc-74_metric_e.html"/>
        <updated>2018-10-08T23:00:00Z</updated>
        <published>2018-10-08T23:00:00Z</published>
        <category term="Weather Forecasts"/>
        <summary type="html">Sunny. High 15. Forecast issued 4:00 PM PDT Monday 08 October 2018</summary>
        <id>tag:weather.gc.ca,2013-04-16:bc-74_fc12:20181008230000</id>
    </entry>
</feed>
`)
}
