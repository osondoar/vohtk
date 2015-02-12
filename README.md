#VoHTK

![asdf](http://i.imgur.com/UIuSRG1.png)

VoHTK is a speech recognition web-application written in Go an powered by a speech recognition engine built with the [HTK](http://htk.eng.cam.ac.uk/) toolkit.

Check it hout [here](http://vohtk.ragar.me).

It uses the new HTML5 [getUserMedia/Stream](http://w3c.github.io/mediacapture-main/getusermedia.html) APIs for recording audio from a compatible browser.

The speech recognition engine is built using the HTK toolkit and based on [Hidden Markov Models](http://en.wikipedia.org/wiki/Hidden_Markov_model). It was originally trained using audio recordings from 630 speakers from the [TIMIT](https://catalog.ldc.upenn.edu/LDC93S1) speech corpus.

This project is based on my Master's Thesis work (_Interactive Map with Voice Control for Mobile Devices_) during my stay at the [Technical University of Liberec](http://www.tul.cz/en/) in 2011.

###License
[MIT License](http://opensource.org/licenses/MIT)
