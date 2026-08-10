const cacheName = self.location.pathname
const pages = [

  "/devtoprod-hugo-blog/blog/20260810181336532/",
  "/devtoprod-hugo-blog/",
  "/devtoprod-hugo-blog/posts/",
  "/devtoprod-hugo-blog/blog/20260810181343536/",
  "/devtoprod-hugo-blog/blog/20260810181410547/",
  "/devtoprod-hugo-blog/blog/20260810130708883/",
  "/devtoprod-hugo-blog/blog/20260810112650010/",
  "/devtoprod-hugo-blog/blog/20260810142237492/",
  "/devtoprod-hugo-blog/blog/20260810102536715/",
  "/devtoprod-hugo-blog/blog/20260810102551730/",
  "/devtoprod-hugo-blog/blog/20260810102652788/",
  "/devtoprod-hugo-blog/blog/20260810000710230/",
  "/devtoprod-hugo-blog/blog/20260810051847844/",
  "/devtoprod-hugo-blog/blog/20260810001047401/",
  "/devtoprod-hugo-blog/blog/20260810003307581/",
  "/devtoprod-hugo-blog/blog/20260810004409703/",
  "/devtoprod-hugo-blog/blog/20260809210137151/",
  "/devtoprod-hugo-blog/blog/20260809214212721/",
  "/devtoprod-hugo-blog/blog/20260809162052915/",
  "/devtoprod-hugo-blog/blog/20260809134442774/",
  "/devtoprod-hugo-blog/blog/20260809135328992/",
  "/devtoprod-hugo-blog/blog/20260809082729705/",
  "/devtoprod-hugo-blog/blog/20260809102017337/",
  "/devtoprod-hugo-blog/blog/20260809052109344/",
  "/devtoprod-hugo-blog/blog/20260809052413348/",
  "/devtoprod-hugo-blog/blog/20260809052209345/",
  "/devtoprod-hugo-blog/blog/20260808185012639/",
  "/devtoprod-hugo-blog/blog/20260808233137182/",
  "/devtoprod-hugo-blog/blog/20260809010909568/",
  "/devtoprod-hugo-blog/blog/20260808201250537/",
  "/devtoprod-hugo-blog/blog/20260808213910352/",
  "/devtoprod-hugo-blog/blog/20260808143240048/",
  "/devtoprod-hugo-blog/blog/20260808172507649/",
  "/devtoprod-hugo-blog/blog/20260808091936858/",
  "/devtoprod-hugo-blog/blog/20260808122250834/",
  "/devtoprod-hugo-blog/blog/20260808134907592/",
  "/devtoprod-hugo-blog/blog/20260808073115760/",
  "/devtoprod-hugo-blog/blog/20260808101107440/",
  "/devtoprod-hugo-blog/blog/20260808052209200/",
  "/devtoprod-hugo-blog/blog/20260807185137728/",
  "/devtoprod-hugo-blog/blog/20260808011107646/",
  "/devtoprod-hugo-blog/blog/20260807215107928/",
  "/devtoprod-hugo-blog/blog/20260807215537992/",
  "/devtoprod-hugo-blog/blog/20260807153409096/",
  "/devtoprod-hugo-blog/blog/20260807100254295/",
  "/devtoprod-hugo-blog/blog/20260807160308168/",
  "/devtoprod-hugo-blog/blog/20260807093117964/",
  "/devtoprod-hugo-blog/blog/20260807143609088/",
  "/devtoprod-hugo-blog/blog/20260807143623095/",
  "/devtoprod-hugo-blog/blog/20260807094650568/",
  "/devtoprod-hugo-blog/blog/20260807093248040/",
  "/devtoprod-hugo-blog/blog/20260807095042739/",
  "/devtoprod-hugo-blog/blog/20260806223927248/",
  "/devtoprod-hugo-blog/blog/20260806192209202/",
  "/devtoprod-hugo-blog/blog/20260806181538775/",
  "/devtoprod-hugo-blog/blog/20260806193332427/",
  "/devtoprod-hugo-blog/blog/20260806150907021/",
  "/devtoprod-hugo-blog/blog/20260806003708893/",
  "/devtoprod-hugo-blog/blog/20260806082908626/",
  "/devtoprod-hugo-blog/blog/20260806111907968/",
  "/devtoprod-hugo-blog/blog/20260806111929989/",
  "/devtoprod-hugo-blog/blog/20260806060151720/",
  "/devtoprod-hugo-blog/blog/20260806060147713/",
  "/devtoprod-hugo-blog/blog/20260806051137370/",
  "/devtoprod-hugo-blog/blog/20260806010609139/",
  "/devtoprod-hugo-blog/blog/20260805225646805/",
  "/devtoprod-hugo-blog/blog/20260805230107874/",
  "/devtoprod-hugo-blog/blog/20260805190250532/",
  "/devtoprod-hugo-blog/blog/20260805152447490/",
  "/devtoprod-hugo-blog/blog/20260805110813042/",
  "/devtoprod-hugo-blog/blog/20260805111107210/",
  "/devtoprod-hugo-blog/blog/20260805055649907/",
  "/devtoprod-hugo-blog/blog/20260805060340197/",
  "/devtoprod-hugo-blog/blog/20260805002850179/",
  "/devtoprod-hugo-blog/blog/20260804210137229/",
  "/devtoprod-hugo-blog/blog/20260804230009212/",
  "/devtoprod-hugo-blog/blog/20260804183610943/",
  "/devtoprod-hugo-blog/blog/20260804190252533/",
  "/devtoprod-hugo-blog/blog/20260804192534011/",
  "/devtoprod-hugo-blog/blog/20260804140254742/",
  "/devtoprod-hugo-blog/blog/20260804152351655/",
  "/devtoprod-hugo-blog/blog/20260804101213904/",
  "/devtoprod-hugo-blog/blog/20260804054111853/",
  "/devtoprod-hugo-blog/blog/20260804111009144/",
  "/devtoprod-hugo-blog/blog/20260804111209244/",
  "/devtoprod-hugo-blog/blog/20260804053730834/",
  "/devtoprod-hugo-blog/blog/20260803214009337/",
  "/devtoprod-hugo-blog/blog/20260803212738199/",
  "/devtoprod-hugo-blog/blog/20260803223137996/",
  "/devtoprod-hugo-blog/blog/20260803172907285/",
  "/devtoprod-hugo-blog/blog/20260803173450548/",
  "/devtoprod-hugo-blog/blog/20260803195137647/",
  "/devtoprod-hugo-blog/blog/20260803155010155/",
  "/devtoprod-hugo-blog/blog/20260803153249329/",
  "/devtoprod-hugo-blog/blog/20260803153137282/",
  "/devtoprod-hugo-blog/blog/20260803110214349/",
  "/devtoprod-hugo-blog/blog/20260803113249908/",
  "/devtoprod-hugo-blog/blog/20260803113255928/",
  "/devtoprod-hugo-blog/blog/20260802212736420/",
  "/devtoprod-hugo-blog/blog/20260803055044301/",
  "/devtoprod-hugo-blog/blog/20260802222928232/",
  "/devtoprod-hugo-blog/blog/20260803000108807/",
  "/devtoprod-hugo-blog/blog/20260802213509660/",
  "/devtoprod-hugo-blog/blog/20260802220907090/",
  "/devtoprod-hugo-blog/blog/20260802181851597/",
  "/devtoprod-hugo-blog/blog/20260802184849117/",
  "/devtoprod-hugo-blog/blog/20260802114937905/",
  "/devtoprod-hugo-blog/blog/20260802151211188/",
  "/devtoprod-hugo-blog/blog/20260802152936551/",
  "/devtoprod-hugo-blog/blog/20260802111809315/",
  "/devtoprod-hugo-blog/blog/20260802111450230/",
  "/devtoprod-hugo-blog/blog/20260802113107582/",
  "/devtoprod-hugo-blog/blog/20260801230107536/",
  "/devtoprod-hugo-blog/blog/20260802053049465/",
  "/devtoprod-hugo-blog/blog/20260801190010150/",
  "/devtoprod-hugo-blog/blog/20260802005241052/",
  "/devtoprod-hugo-blog/blog/20260801215651055/",
  "/devtoprod-hugo-blog/blog/20260801220415131/",
  "/devtoprod-hugo-blog/blog/20260801221247247/",
  "/devtoprod-hugo-blog/blog/20260801154536064/",
  "/devtoprod-hugo-blog/blog/20260801071652723/",
  "/devtoprod-hugo-blog/blog/20260801184843089/",
  "/devtoprod-hugo-blog/blog/20260801133551705/",
  "/devtoprod-hugo-blog/blog/20260801150107539/",
  "/devtoprod-hugo-blog/blog/20260801112330529/",
  "/devtoprod-hugo-blog/blog/20260801090120718/",
  "/devtoprod-hugo-blog/blog/20260801111851483/",
  "/devtoprod-hugo-blog/blog/20260801113209641/",
  "/devtoprod-hugo-blog/blog/20260801053210700/",
  "/devtoprod-hugo-blog/blog/20260731230209761/",
  "/devtoprod-hugo-blog/blog/20260731214652456/",
  "/devtoprod-hugo-blog/blog/20260731225736710/",
  "/devtoprod-hugo-blog/blog/20260731225611679/",
  "/devtoprod-hugo-blog/blog/20260731184708132/",
  "/devtoprod-hugo-blog/blog/20260731190209579/",
  "/devtoprod-hugo-blog/blog/20260731153336536/",
  "/devtoprod-hugo-blog/blog/20260731113209490/",
  "/devtoprod-hugo-blog/blog/20260731072102597/",
  "/devtoprod-hugo-blog/blog/20260731111707859/",
  "/devtoprod-hugo-blog/blog/20260731060336361/",
  "/devtoprod-hugo-blog/blog/20260731010851587/",
  "/devtoprod-hugo-blog/blog/20260731013447793/",
  "/devtoprod-hugo-blog/blog/20260730223050128/",
  "/devtoprod-hugo-blog/blog/20260730224809353/",
  "/devtoprod-hugo-blog/blog/20260730180143772/",
  "/devtoprod-hugo-blog/blog/20260730190927525/",
  "/devtoprod-hugo-blog/blog/20260730191515714/",
  "/devtoprod-hugo-blog/blog/20260730143243571/",
  "/devtoprod-hugo-blog/blog/20260730104209782/",
  "/devtoprod-hugo-blog/blog/20260730152451676/",
  "/devtoprod-hugo-blog/blog/20260730152410649/",
  "/devtoprod-hugo-blog/blog/20260730105209309/",
  "/devtoprod-hugo-blog/blog/20260730030128343/",
  "/devtoprod-hugo-blog/blog/20260730012928049/",
  "/devtoprod-hugo-blog/blog/20260729225649097/",
  "/devtoprod-hugo-blog/blog/20260729225910140/",
  "/devtoprod-hugo-blog/blog/20260729183409297/",
  "/devtoprod-hugo-blog/blog/20260729192614538/",
  "/devtoprod-hugo-blog/blog/20260729193127654/",
  "/devtoprod-hugo-blog/blog/20260729152611635/",
  "/devtoprod-hugo-blog/blog/20260729111912723/",
  "/devtoprod-hugo-blog/blog/20260729105137109/",
  "/devtoprod-hugo-blog/blog/20260729112307909/",
  "/devtoprod-hugo-blog/blog/20260728232708840/",
  "/devtoprod-hugo-blog/blog/20260729060323461/",
  "/devtoprod-hugo-blog/blog/20260729010508970/",
  "/devtoprod-hugo-blog/blog/20260729010128940/",
  "/devtoprod-hugo-blog/blog/20260729011810069/",
  "/devtoprod-hugo-blog/blog/20260728225313581/",
  "/devtoprod-hugo-blog/blog/20260728225938625/",
  "/devtoprod-hugo-blog/blog/project-introduction/",
  "/devtoprod-hugo-blog/blog/20260728181249835/",
  "/devtoprod-hugo-blog/blog/20260728192754806/",
  "/devtoprod-hugo-blog/blog/20260728153109335/",
  "/devtoprod-hugo-blog/blog/20260728170337861/",
  "/devtoprod-hugo-blog/blog/20260728170855091/",
  "/devtoprod-hugo-blog/blog/20260728150307881/",
  "/devtoprod-hugo-blog/blog/20260728111948829/",
  "/devtoprod-hugo-blog/blog/20260728103334347/",
  "/devtoprod-hugo-blog/blog/20260728170610973/",
  "/devtoprod-hugo-blog/blog/book-sample/",
  "/devtoprod-hugo-blog/blog/dev-sample/",
  "/devtoprod-hugo-blog/categories/",
  "/devtoprod-hugo-blog/list/",
  "/devtoprod-hugo-blog/list/newest/",
  "/devtoprod-hugo-blog/list/oldest/",
  "/devtoprod-hugo-blog/tags/",
  "/devtoprod-hugo-blog/search/",
  "/devtoprod-hugo-blog/main.min.f9c55459aa1fb5a9ce5b6ae28b563e5f91180c6a8549b7a74acdcf81e3977f28.css",
  
];

self.addEventListener("install", function (event) {
  self.skipWaiting();

  caches.open(cacheName).then((cache) => {
    return cache.addAll(pages);
  });
});

self.addEventListener("fetch", (event) => {
  const request = event.request;
  if (request.method !== "GET") {
    return;
  }

  /**
   * @param {Response} response
   * @returns {Promise<Response>}
   */
  function saveToCache(response) {
    if (cacheable(response)) {
      return caches
        .open(cacheName)
        .then((cache) => cache.put(request, response.clone()))
        .then(() => response);
    } else {
      return response;
    }
  }

  /**
   * @param {Error} error
   */
  function serveFromCache(error) {
    return caches.open(cacheName).then((cache) => cache.match(request.url));
  }

  /**
   * @param {Response} response
   * @returns {Boolean}
   */
  function cacheable(response) {
    return response.type === "basic" && response.ok && !response.headers.has("Content-Disposition")
  }

  event.respondWith(fetch(request).then(saveToCache).catch(serveFromCache));
});