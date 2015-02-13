jQuery(function($) {

  var recorder;
  var audioCtx = new (window.AudioContext || window.webkitAudioContext)();
  audioCtx.sampleRate = 16000
  var consecutiveSilenceCycles = 0;
  var soundCycles = 0;
  var recordingFinished = false;

  checkUserMedia();

  $("#microphone").click(toggleRecording);

  function checkUserMedia(){
    navigator.getUserMedia = (navigator.getUserMedia ||
    navigator.webkitGetUserMedia ||
    navigator.mozGetUserMedia ||
    navigator.msGetUserMedia);



    if (navigator.getUserMedia) {
      console.log('getUserMedia is supported!');
    } else {
      alert('Stream API is not supported in your browser');
      console.log('getUserMedia not supported on your browser!');
    }
  }

  function initializeAutoRecording(){
    consecutiveSilenceCycles = 0;
    soundCycles = 0;
    recordingFinished = false;
  }

  function recordingCallback(data){
    var total = 0;
    var amplitude;
    for(var i = 0; i < data.length; i++){
      total += Math.abs(data[i]);
    }
    amplitudeAverage = Math.sqrt( total / data.length )
    graphicSize = 80 + ( 200 * amplitudeAverage );
    $(".recording-animation .circle").css( {"width": graphicSize, "height": graphicSize} );
    if ( $('input[id=auto-stop]:checked').attr('id') ){
      autoStop();
    }
    return amplitude;
  }

  function autoStop(averageAmplitude){
    if(amplitudeAverage < 0.25){
      consecutiveSilenceCycles+= 1;
      if(soundCycles > 3 && consecutiveSilenceCycles > 8) {
        recordingFinished  = true;
        stopRecording();
      }
    }
    else{
      recordingFinished  = false;
      soundCycles+= 1 ;
      consecutiveSilenceCycles = 0;
    }

    console.log(amplitudeAverage, consecutiveSilenceCycles, soundCycles ,recordingFinished);

  }

  function toggleRecording(e){
    if($(this).hasClass("active-microphone")){
      stopRecording();
    }
    else{
      startRecording();
    }
  }


  function stopRecording(){
    $("#microphone").toggleClass("active-microphone");
    $(".recording-animation").css("visibility","hidden");
    $(".cd-intro").css("display","block");
    $("#transcribed-text-wrapper").css("display","block");
    recorder.stop();
    recorder.exportWAV(function(blob){
      sendRawAudio(blob);
    });
  }

  function startRecording(inputBuffer){
    initializeAutoRecording();
    navigator.getUserMedia(
      { audio: true },
      // Success callback
      function(stream) {
        var source = audioCtx.createMediaStreamSource(stream);
        recorder = new Recorder(source, { numChannels: 1, sampleRate: 16000, recordingCallback: recordingCallback });
        recorder.record();
        $("#microphone").toggleClass("active-microphone");
        $(".recording-animation").css("visibility","visible");
        $(".cd-intro").css("display","none");
        $("#transcribed-text-wrapper").css("display","none");
      },
      function(err) {
        console.log('The following gUM error occured: ' , err);
      }
    )
  }


  function sendRawAudio(data){
    var fileType = 'raw'; // or "audio"
    var fileName = 'audio.raw';  // or "wav" or "ogg"

    var blob = data
    var formData = new FormData();
    // formData.append(fileType + '-filename', fileName);
    formData.append(fileType + '-blob', blob);

    $.ajax( {
      url: "api/requests",
      type: "POST",
      data: formData,
      processData: false,  // tell jQuery not to process the data
      contentType: false,   // tell jQuery not to set contentType
      success: recognitionCallback,
      error: recognitionErrorCallback
    });

  }

  function recognitionCallback(response){
    $("#transcribed-text h3").html(response.transcription)
  }

  function recognitionErrorCallback(error){
    console.log(error);
  }


  var utteranceExamples = [
    "Zoom in",
    "Search for the closest hotel",
    "Where are the most expensive cinemas in granada",
    "Find the cheapest train station",
    "Search for the cheapest chinese restaurant",
    "Search for hotels",
    "Go to New York",
    "Directions from Prague to chicago",
    "Find the best vegetarian restaurant",
    "Search for the cheapest church in Detroit",
    "Directions from W to chicago",
    "Where is the cheapest vegetarian restaurant in Detroit",
    "Where are the closest tram stops in liberec",
    "Where is the closest chinese restaurant in Prague",
    "Directions from Spain to germany",
    "Move left",
    "Directions from New York to Washington",
    "Where is the closest japanese restaurant in Italy",
    "View in satellite mode",
    "Search for expensive restaurants in germany",
    "Find restaurants",
    "Where is the best cinema in Spain",
    "Go to dallas",
    "Find the best vegetarian restaurant",
    "Find hospitals",
    "Search for churches in chicago",
    "Where are the cheapest hospitals",
    "Where am ",
    "Move left",
    "Find chinese restaurants",
    "Directions from Spain to Detroit",
    "Set zoom level to sixteen",
    "Set zoom level to ten",
    "View in satellite mode",
    "Look for vegetarian restaurants",
    "Where is the most expensive theater in Prague",
    "Look for close spanish restaurants",
    "Find nearby metro stations",
    "Set zoom level to four",
    "Directions from Czech Republic to huelma",
    "Search for spanish restaurants",
    "Where is the nearest hospital",
    "Directions from liberec to Prague",
    "Go to france",
    "Set zoom level to twenty",
    "Move right",
    "Where is the closest cinema",
    "Search for the best hospital",
    "Look for hospitals",
    "Set zoom level to eighteen",
    "Set zoom level to max",
    "Move down",
    "Zoom out",
    "Set zoom level to max",
    "Where are the closest hotels in huelma",
    "Search for the cheapest vegetarian restaurant in Washington",
    "Look for the most expensive tram stop",
    "Find the cheapest church",
    "Move left",
    "Zoom out",
    "Directions from Italy to Washington",
    "Find the most expensive theater in Czech Republic",
    "Find churches",
    "Directions from Spain to here",
    "Directions from Detroit to Prague",
    "Look for japanese restaurants in germany",
    "Look for the most expensive chinese restaurant",
    "Look for cheap vegetarian restaurants",
    "Look for cheap hospitals in liberec",
    "Look for the nearest vegetarian restaurant",
    "Look for metro stations in New York",
    "Find the cheapest church",
    "Move left",
    "Search for pharmacies in france",
    "Find nearby metro stations in madrid",
    "Set zoom level to four",
    "Set zoom level to sixteen",
    "Where are the cheapest train stations in liberec",
    "Where is the nearest hospital",
    "Go to chicago",
    "Search for the nearest restaurant in madrid",
    "Set zoom level to zero",
    "Find tram stops",
    "Search for the closest mexican restaurant",
    "Set zoom level to one",
    "Find the closest metro station",
    "Where are the cheapest chinese restaurants",
    "Look for metro stations in Spain",
    "Go to liberec",
    "Search for pharmacies",
    "Move left",
    "Look for restaurants in New York",
    "Look for the most expensive vegetarian restaurant in Detroit",
    "Find the most expensive cinema",
    "Find chinese restaurants in Prague",
    "Where is the cheapest theater in huelma",
    "Look for the cheapest train station in New York",
    "Where are the most expensive restaurants",
    "Move up",
    "Search for japanese restaurants in Czech Republic",
    "Go to Italy",
    "Find the best spanish restaurant",
    "Set zoom level to thirteen",
    "Go to Washington",
    "Look for tram stops",
    "Find pharmacies",
    "Directions from Italy to france",
    "Move up",
    "Search for the most expensive tram stop in Italy",
    "Find the nearest hotel in Prague",
    "Go to Prague",
    "Directions from granada to huelma",
    "Find close churches in madrid",
    "Find nearby metro stations in liberec",
    "Where are the nearest spanish restaurants in Detroit",
    "View in satellite mode",
    "Where is the nearest church in france",
    "Move right",
    "Where is the closest tram stop in madrid",
    "Where are the best japanese restaurants",
    "Directions from Italy to dallas",
    "Set zoom level to maximum",
    "Find the most expensive chinese restaurant",
    "Where is the most expensive train station",
    "Where are the best restaurants in New York",
    "View in satellite mode",
    "Search for nearby tram stops in liberec",
    "Search for the most expensive theater in Italy",
    "View in satellite mode",
    "Look for tram stops in Washington",
    "Find theaters",
    "Where is the closest japanese restaurant in Prague",
    "Find the best church",
    "Where are the best japanese restaurants",
    "Search for the most expensive chinese restaurant",
    "Search for metro stations",
    "Look for vegetarian restaurants in madrid",
    "Go to Czech Republic"
  ];
  function displayExamples(){

    var examples = $('.cd-words-wrapper b');
    for(var j, i = utteranceExamples.length; i; i--){
      j = Math.floor(Math.random() * i),
      e1 = $('.cd-words-wrapper b:nth-child(' + i + ')');
      e2 = $('.cd-words-wrapper b:nth-child(' + j + ')');
      x = e1.html();;
      e1.html(e2.html());
      e2.html(x);
    }


    // $('.cd-words-wrapper').append('<b class="is-visible">' + utteranceExamples[i] + '</b>');
    // for(var i=1; i < utteranceExamples.length; i++) {
    //   console.log("<b>" + utteranceExamples[i] + "</b>");
    //   $('.cd-words-wrapper').append("<b>" + utteranceExamples[i] + "</b>");
    // }
  }
  displayExamples();


});
