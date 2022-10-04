import React, { Component } from "react";
import "./Camposorte.css";

class Camposorte extends Component {
    constructor(props) {
        super(props);
        this.state = {
            msg: this.props.msg,
            anim: false,
        };

        // this.test = this.test.bind(this);
    }

    componentDidUpdate() {
        // eslint-disable-next-line
        if (this.props.msg != this.state.msg) {
            this.setState({ anim: true });
            this.setState({ msg: this.props.msg });
        }
    }

    // test() {
    //     this.setState({ msg: this.props.msg });
    // }

    render() {
        return (
            <div
                onAnimationEnd={() => {
                    this.setState({ anim: false });
                    this.setState({ msg: this.props.msg });
                }}
                id="result"
                className={this.state.anim ? "anim" : ""}
            >
                <p>{!this.state.anim && this.state.msg}</p>
            </div>
        );
    }
}

export default Camposorte;
