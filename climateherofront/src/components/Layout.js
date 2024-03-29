import React, { Component } from 'react';
import { Col, Grid, Row } from 'react-bootstrap';

export class Layout extends Component {
  displayName = Layout.name

  render() {
    return (
        <Row>
          <Col sm={9}>
            {this.props.children}
          </Col>
        </Row>
    );
  }
}