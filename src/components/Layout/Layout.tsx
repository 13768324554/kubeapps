import * as React from "react";

import Footer from "../Footer";
import Header from "../Header";
import Sidebar from "../Sidebar";
import "./Layout.css";

class Layout extends React.Component {
  public render() {
    return (
      <section className="Layout">
        <Header />
        <main>
          <div className="container container-fluid padding-reset">
            <div className="row">
              <div className="col-1">
                <Sidebar />
              </div>
              <div className="col-11">{this.props.children}</div>
            </div>
          </div>
        </main>
        <Footer />
      </section>
    );
  }
}

export default Layout;
