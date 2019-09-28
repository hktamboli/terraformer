package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ExpressCloudConnectionType is a nested struct in vpc response
type ExpressCloudConnectionType struct {
	InstanceId                    string                    `json:"InstanceId" xml:"InstanceId"`
	Status                        string                    `json:"Status" xml:"Status"`
	Name                          string                    `json:"Name" xml:"Name"`
	Description                   string                    `json:"Description" xml:"Description"`
	GmtCreate                     string                    `json:"GmtCreate" xml:"GmtCreate"`
	GmtModify                     string                    `json:"GmtModify" xml:"GmtModify"`
	PeerCity                      string                    `json:"PeerCity" xml:"PeerCity"`
	PeerLocation                  string                    `json:"PeerLocation" xml:"PeerLocation"`
	PortType                      string                    `json:"PortType" xml:"PortType"`
	Bandwidth                     int                       `json:"Bandwidth" xml:"Bandwidth"`
	Distance                      int                       `json:"Distance" xml:"Distance"`
	RedundantEccId                string                    `json:"RedundantEccId" xml:"RedundantEccId"`
	CircuitCode                   string                    `json:"CircuitCode" xml:"CircuitCode"`
	Isp                           string                    `json:"Isp" xml:"Isp"`
	Type                          string                    `json:"Type" xml:"Type"`
	IdcSP                         string                    `json:"IdcSP" xml:"IdcSP"`
	BusinessStatus                string                    `json:"BusinessStatus" xml:"BusinessStatus"`
	HasReservationData            string                    `json:"HasReservationData" xml:"HasReservationData"`
	ReservationBandwidth          string                    `json:"ReservationBandwidth" xml:"ReservationBandwidth"`
	ReservationInternetChargeType string                    `json:"ReservationInternetChargeType" xml:"ReservationInternetChargeType"`
	ReservationActiveTime         string                    `json:"ReservationActiveTime" xml:"ReservationActiveTime"`
	ReservationOrderType          string                    `json:"ReservationOrderType" xml:"ReservationOrderType"`
	ApplicationType               string                    `json:"ApplicationType" xml:"ApplicationType"`
	ApplicationId                 string                    `json:"ApplicationId" xml:"ApplicationId"`
	ApplicationStatus             string                    `json:"ApplicationStatus" xml:"ApplicationStatus"`
	ApplicationBandwidth          string                    `json:"ApplicationBandwidth" xml:"ApplicationBandwidth"`
	ContactTel                    string                    `json:"ContactTel" xml:"ContactTel"`
	ContactMail                   string                    `json:"ContactMail" xml:"ContactMail"`
	IDCardNo                      string                    `json:"IDCardNo" xml:"IDCardNo"`
	EndTime                       string                    `json:"EndTime" xml:"EndTime"`
	ChargeType                    string                    `json:"ChargeType" xml:"ChargeType"`
	VirtualBorderRouterModels     VirtualBorderRouterModels `json:"VirtualBorderRouterModels" xml:"VirtualBorderRouterModels"`
}
