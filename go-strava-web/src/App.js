import React, { useState } from 'react';
import "../node_modules/gestalt/dist/gestalt.css"
import { Box, Button, Container, Text } from "gestalt";
import { Circle, CircleMarker, Map, Marker, Polygon, Polyline, Popup, Rectangle, TileLayer, Tooltip } from "react-leaflet"

import queryString from 'querystring'

const center = [51.505, -0.09]

const multiPolygon = [
    [
        [51.51, -0.12],
        [51.51, -0.13],
        [51.53, -0.13],
    ],
    [
        [51.51, -0.05],
        [51.51, -0.07],
        [51.53, -0.07],
    ],
]

const rectangle = [
    [51.49, -0.08],
    [51.5, -0.06],
]

function App() {
    console.log(window.location.search);
    const parsed = queryString.parse(window.location.search);
    console.log(parsed);

    const handleClick = (e) => {
        const client_id = 50658;
        const redirect_uri = "http://localhost:3000"
        const response_type = 'code'
        const scope = 'activity:read'
        let url = `https://www.strava.com/oauth/authorize?client_id=${client_id}&redirect_uri=${redirect_uri}&response_type=${response_type}&scope=${scope}`
        window.location.assign(url);
    }

  return (
      <Box color="gray" padding={3}>
          <Container>
              <Box color="white" padding={3}>
                  <Text>Centered content</Text>
                  <Button onClick={handleClick} text="Login" inline />
              </Box>
          </Container>
          <Container>
              <Map center={center} zoom={18}>
                  <TileLayer
                      url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                  />
                  <Circle
                      center={center}
                      fillColor="blue"
                      radius={200}>
                  </Circle>
                  <CircleMarker center={[51.51, -0.12]} color="red" radius={20}>
                      <Tooltip>Tooltip for CircleMarker</Tooltip>
                  </CircleMarker>
                  <Marker position={[51.51, -0.09]}>
                      <Popup>Popup for Marker</Popup>
                      <Tooltip>Tooltip for Marker</Tooltip>
                  </Marker>
                  <Polygon color="purple" positions={multiPolygon}>
                      <Tooltip sticky>sticky Tooltip for Polygon</Tooltip>
                  </Polygon>
                  <Rectangle bounds={rectangle} color="black">
                      <Tooltip direction="bottom" offset={[0, 20]} opacity={1} permanent>
                          permanent Tooltip for Rectangle
                      </Tooltip>
                  </Rectangle>
              </Map>
          </Container>
      </Box>
  );
}

export default App;
