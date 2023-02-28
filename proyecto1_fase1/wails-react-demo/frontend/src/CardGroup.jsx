import React from 'react';
import {
  MDBCard,
  MDBCardImage,
  MDBCardBody,
  MDBCardTitle,
  MDBCardText,
  MDBRow,
  MDBCol
} from 'mdb-react-ui-kit';

import disk from './assets/images/disk.jpg';
import cpu from './assets/images/cpu.jpg';


export default function CardGroup() {
  return (
    <MDBRow className='row-cols-1 row-cols-md-2 g-2' style={{ paddingLeft: 25, paddingRight: 25, marginTop: -75 }}>
      <MDBCol>
        <MDBCard className='h-100'>
          <MDBCardImage
            src={disk}
            alt='...'
            style={{ height: 165 }}
            position='top'
          />
          <MDBCardBody>
            <MDBCardTitle>Disco duro</MDBCardTitle>
            <MDBCardText>
              d
            </MDBCardText>
          </MDBCardBody>
        </MDBCard>
      </MDBCol>
      <MDBCol>
        <MDBCard className='h-100'>
          <MDBCardImage
            src={cpu}
            alt='...'
            style={{ height: 165 }}
            position='top'
          />
          <MDBCardBody>
            <MDBCardTitle>CPU</MDBCardTitle>
            <MDBCardText>
              T
            </MDBCardText>
          </MDBCardBody>
        </MDBCard>
      </MDBCol>
    </MDBRow>
  );
}