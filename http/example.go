package http

import (
	"fmt"
	"net/http"
)
const htmlCode = `
<!-- LAYOUT CONTAINER -->
<!-- Simple header with fixed tabs. -->
<div class="mdl-layout mdl-js-layout mdl-layout--fixed-header
            mdl-layout--fixed-tabs">
  <header class="mdl-layout__header">
    <div class="mdl-layout__header-row">
      <!-- Title -->
      <span class="mdl-layout-title">Kaptain Kitty</span>
    </div>
    <!-- Tabs -->
    <div class="mdl-layout__tab-bar mdl-js-ripple-effect">
      <a href="#fixed-tab-1" class="mdl-layout__tab is-active">About</a>
      <a href="#fixed-tab-2" class="mdl-layout__tab">Learn</a>
    </div>
  </header>

  <div class="mdl-layout__drawer">
    <span class="mdl-layout-title">Kaptain Kitty</span>
    <div class="avatar">
      <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/234228/cat.jpg" alt="Kaptain Kitty" class="avatar-img">
    </div>
    <!-- /.avatar -->

    <div class="drawer-text">
      Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aspernatur officiis animi, soluta ab deserunt dolore fugit voluptatem laboriosam, magni. Eligendi quia quasi qui cupiditate optio fugit vel, suscipit harum illum.
    </div>
    <!-- /.drawer-text -->
  </div>
  <!-- /.mdl-layout__drawer -->

  <main class="mdl-layout__content">

    <div class="mdl-layout__tab-panel is-active" id="fixed-tab-1">
      <div class="page-content">
        <!-- Your content goes here -->
        <!-- Hero section -->
        <div class="hero-section">

          <div class="hero-text mdl-typography--text-center">

            <h1 class="mdl-typography--display-2">I'm Kaptain Kitty</h1>
            <p class="mdl-typography--display-1">
              I'll teach you loads about your kitten
            </p>

            <a class="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect mdl-button--accent kitty-hero__text-button" href="#intro">
              <i class="material-icons">keyboard_arrow_down</i>
            </a>

          </div>
          <!-- /.hero-text -->

        </div>
        <!-- /.hero-section -->

        <!-- INTRO -->
        <div id="intro" class="mdl-grid intro-section">
          <div class="about-kitty mdl-cell mdl-cell--12-col">
            <p class="mdl-typography--headline">
              Welcome to Kaptain Kitty!
              This is a demo HTML template that accompanies an article for SitePoint. The article illustrates how to use the Material Design Lite library to build a web page. You're free to use this template as you like. All images on this demo are courtesy of <a href="https://pixabay.com/">Pixabay.com</a>.
            </p>
          </div>
          <!-- /.about-kitty -->

          <div class="about-kitty mdl-cell mdl-cell--12-col">
            <p>
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Sint dolorum consectetur natus sequi, est similique! Temporibus rem consequuntur laudantium, illo excepturi velit quas. Culpa ipsum dolor tempore accusantium sed iusto.
            </p>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ut, dicta aliquid, inventore a ullam excepturi similique sit, nobis incidunt laboriosam amet maxime iusto nam facilis possimus quo optio? Dolor, et?
            </p>
          </div>
          <!-- /.about-kitty -->

          <div class="about-kitty mdl-cell mdl-cell--5-col mdl-cell--1-col-tablet mdl-cell--hide-phone">
            <div class="circle-container">
              <div class="circle"></div>
              <div class="circle"></div>
              <div class="circle"></div>
              <div class="circle"></div>
              <div class="circle"></div>
            </div>
            <!-- /.circle-container -->

          </div>
          <!-- /.about-kitty -->
          <div class="about-kitty mdl-cell mdl-cell--7-col mdl-cell--6-col-tablet mdl-cell--4-col-phone">
            <div class="topics-container">
              <div class="topic">Feeding</div>
              <div class="topic">Choosing the right vet</div>
              <div class="topic">Keeping you kittens healthy</div>
              <div class="topic">Adopting a kitten</div>
              <div class="topic">Vaccinating your kitten</div>
            </div>
            <!-- /.topics-container -->
          </div>
          <!-- /.about-kitty -->

          <div class="about-kitty mdl-cell mdl-cell--12-col">

            <p class="clearfix">
              <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/234228/cat.jpg" alt="Kaptain Kitty" class="embedded-img float-left"> Lorem ipsum dolor sit amet, consectetur adipisicing elit. Sint dolorum consectetur natus sequi, est similique! Temporibus
              rem consequuntur laudantium, illo excepturi velit quas. Culpa ipsum dolor tempore accusantium sed iusto. Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ut, dicta aliquid, inventore a ullam excepturi similique sit, nobis incidunt
              laboriosam amet maxime iusto nam facilis possimus quo optio? Dolor, et?
            </p>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ut, dicta aliquid, inventore a ullam excepturi similique sit, nobis incidunt laboriosam amet maxime iusto nam facilis possimus quo optio? Dolor, et?
            </p>
          </div>
          <!-- /.about-kitty -->
        </div>
        <!--/.mdl-grid -->

        <!-- Testimonial -->
        <div class="mdl-grid mdl-grid--no-spacing fullwidth-panel">
          <div class="mdl-cell mdl-cell--12-col mdl-typography--text-center quote-panel">
            <blockquote>
              <p>
                Lorem ipsum dolor sit amet, consectetur adipisicing elit. Sint dolorum consectetur natus sequi, est similique! Temporibus rem consequuntur laudantium, illo excepturi velit quas. Culpa ipsum dolor tempore accusantium sed iusto.
              </p>
              <footer>
                — <cite>Happy Kitten Owner</cite>
              </footer>
            </blockquote>

          </div>
          <!-- /.mdl-cell -->
        </div>
        <!--/.mdl-grid -->

      </div>
      <!-- /.page-content -->
    </div>
    <!-- /.tab1 -->

    <div class="mdl-layout__tab-panel" id="fixed-tab-2">
      <div class="page-content">
        <!-- Your content goes here -->

        <!-- CARDS  -->
        <div class="mdl-grid cards-section">
          <div class="mdl-cell mdl-cell--6-col mdl-cell--12-col-tablet mdl-card mdl-shadow--2dp home-bringing-card">
            <div class="mdl-card__title">
              <h2 class="mdl-card__title-text">Bringing a kitten into your home</h2>
            </div>
            <div class="mdl-card__supporting-text">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo reiciendis corporis, optio animi autem, quisquam, temporibus sed deleniti officiis dolor eveniet tempore ad nulla deserunt laudantium, obcaecati esse! Officia.
            </div>
          </div>
          <!-- /.mdl-card -->
          <div class="mdl-cell mdl-cell--4-col mdl-cell--4-col-tablet mdl-cell--4-col-phone mdl-card mdl-shadow--2dp play-card">
            <div class="mdl-card__title">
              <h2 class="mdl-card__title-text">Playing with your kitten</h2>
            </div>
            <div class="mdl-card__supporting-text">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo reiciendis corporis, optio animi autem, quisquam, temporibus sed deleniti officiis dolor eveniet tempore ad nulla deserunt laudantium, obcaecati esse! Officia.
            </div>
          </div>
          <!-- /.mdl-card -->
          <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-tablet mdl-cell--4-col-phone mdl-card mdl-shadow--2dp image-card">
            <div class="mdl-card__actions">
              <span class="image-card__title">KaptainKitty.jpg</span>
            </div>
          </div>
          <!-- /.mdl-card -->

          <div class="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet mdl-cell--4-col-phone mdl-card mdl-shadow--2dp litter-card">
            <div class="mdl-card__title">
              <h2 class="mdl-card__title-text">Taking care of a litter of kittens</h2>
            </div>
            <div class="mdl-card__supporting-text">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo reiciendis corporis, optio animi autem, quisquam, temporibus sed deleniti officiis dolor eveniet tempore ad nulla deserunt laudantium, obcaecati esse! Officia.
            </div>
          </div>
          <!--/.mdl-card -->
          <div class="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet mdl-cell--4-col-phone mdl-card mdl-shadow--2dp diet-card">
            <div class="mdl-card__title">
              <h2 class="mdl-card__title-text">Healthy diet for your kitten</h2>
            </div>
            <div class="mdl-card__supporting-text">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo reiciendis corporis, optio animi autem, quisquam, temporibus sed deleniti officiis dolor eveniet tempore ad nulla deserunt laudantium, obcaecati esse! Officia.
            </div>
            <div class="mdl-card__actions mdl-card--border">
              <a class="mdl-button mdl-button--accent mdl-js-button mdl-js-ripple-effect">
									Learn about feeding
							</a>
            </div>
          </div>
          <!--/.mdl-card -->

          <div class="mdl-cell mdl-cell--3-col mdl-cell--8-col-tablet mdl-cell--4-col-phone mdl-card mdl-shadow--2dp card-small">
            <div class="mdl-card__title">
              <h2 class="mdl-card__title-text">Kitten behavior</h2>
            </div>
            <div class="mdl-card__supporting-text">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo reiciendis corporis, optio animi autem, quisquam, temporibus sed deleniti officiis dolor eveniet tempore ad nulla deserunt laudantium, obcaecati esse! Officia.
            </div>
            <div class="mdl-card__actions mdl-card--border">
              <a class="mdl-button mdl-button--accent mdl-js-button mdl-js-ripple-effect">
									More on behavior
							</a>
            </div>
          </div>
          <!--/.mdl-card -->
          <div class="mdl-cell mdl-cell--3-col mdl-cell--8-col-tablet mdl-cell--4-col-phone mdl-card mdl-shadow--2dp card-small">
            <div class="mdl-card__title">
              <h2 class="mdl-card__title-text">Finding a vet</h2>
            </div>
            <div class="mdl-card__supporting-text">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo reiciendis corporis, optio animi autem, quisquam, temporibus sed deleniti officiis dolor eveniet tempore ad nulla deserunt laudantium, obcaecati esse! Officia.
            </div>
            <div class="mdl-card__actions mdl-card--border">
              <a class="mdl-button mdl-button--accent mdl-js-button mdl-js-ripple-effect">
									More on finding a vet
							</a>
            </div>
          </div>
          <!--/.mdl-card -->
          <div class="mdl-cell mdl-cell--3-col mdl-cell--8-col-tablet mdl-cell--4-col-phone mdl-card mdl-shadow--2dp card-small">
            <div class="mdl-card__title">
              <h2 class="mdl-card__title-text">Toys for kittens</h2>
            </div>
            <div class="mdl-card__supporting-text">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo reiciendis corporis, optio animi autem, quisquam, temporibus sed deleniti officiis dolor eveniet tempore ad nulla deserunt laudantium, obcaecati esse! Officia.
            </div>
            <div class="mdl-card__actions mdl-card--border">
              <a class="mdl-button mdl-button--accent mdl-js-button mdl-js-ripple-effect">
									More on toys
							</a>
            </div>
          </div>
          <!--/.mdl-card -->
          <div class="mdl-cell mdl-cell--3-col mdl-cell--8-col-tablet mdl-cell--4-col-phone mdl-card mdl-shadow--2dp card-small">
            <div class="mdl-card__title">
              <h2 class="mdl-card__title-text">Adopting a kitten</h2>
            </div>
            <div class="mdl-card__supporting-text">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolorem explicabo reiciendis corporis, optio animi autem, quisquam, temporibus sed deleniti officiis dolor eveniet tempore ad nulla deserunt laudantium, obcaecati esse! Officia.
            </div>
            <div class="mdl-card__actions mdl-card--border">
              <a class="mdl-button mdl-button--accent mdl-js-button mdl-js-ripple-effect">
									More on adopting
							</a>
            </div>
          </div>
          <!--/.mdl-card -->
        </div>
        <!--/.mdl-grid -->

      </div>
      <!-- /.page-content -->
    </div>
    <!-- /tab2 -->

    <!-- Contact -->
    <div class="mdl-grid mdl-grid--no-spacing">

      <div class="mdl-cell mdl-cell--12-col contact-intro mdl-color--indigo-900">
        <h2 class="mdl-typography--title mdl-typography--title-color-contrast mdl-typography--font-thin mdl-typography--text-center">Make your kittens happier with my tips</h2>
      </div>
      <!--/.contact-intro  -->

      <div class="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet mdl-cell--4-col-phone contact-panel form-panel mdl-color--indigo-50">

        <form action="#">
          <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
            <input class="mdl-textfield__input" type="text" id="name">
            <label class="mdl-textfield__label" for="name">Your name</label>
          </div>
          <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
            <input class="mdl-textfield__input" type="email" id="email">
            <label class="mdl-textfield__label" for="email">Your email</label>
          </div>
          <div class="button-container clearfix">
            <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent subscribe-button">
              Join my fans
            </button>
          </div>
          <!--/.button-container -->

        </form>
      </div>
      <!--/.contact-panel -->

      <div class="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet mdl-cell--4-col-phone contact-panel address-panel mdl-typography--text-center mdl-color--indigo-100">

        <p class="mdl-typography--title-color-contrast mdl-typography--text-nowrap mdl-typography--font-thin">
          <i class="material-icons">email</i> <a href="mailto:info@captain.kitty.com">info@kaptain.kitty.com</a>
        </p>

        <p class="mdl-typography--title-color-contrast mdl-typography--text-nowrap mdl-typography--font-thin">
          <a class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect" href="twitter.com">twitter</a>
          <a class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect" href="plus.google.com">Google+</a>
          <a class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect" href="facebook.com">Facebook</a>
        </p>
      </div>
      <!-- /.contact-panel -->

    </div>
    <!-- /.mdl-grid -->

    <!-- FOOTER -->
    <footer class="mdl-mini-footer mdl-color--indigo-200">
      <div class="mdl-mini-footer__left-section">
        <div class="mdl-logo">Kaptain Kitty &ndash; designed by <a href="http://wpthememakeover.com">Maria Antonietta Perna</a></div>
        <ul class="mdl-mini-footer__link-list">
          <li><a href="#">Help</a></li>
          <li><a href="#">Privacy & Terms</a></li>
        </ul>
        &copy; Maria Antonietta Perna 2016
      </div>
      <!-- /.mdl-mini-footer__left-section -->
    </footer>

  </main>
  <!-- /.mdl-layout__content -->
</div>
<!-- /.mdl-layout -->
`
func home(res http.ResponseWriter, req *http.Request){
url := req.FormValue("url")
if url == ""{
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(res, htmlCode)
	return 
}
}
func Handle(){
	http.HandleFunc("/home", home)
	http.ListenAndServe(":3000", nil)
}