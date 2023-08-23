package v1

import (
	"bytes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/web/adapters"
	"github.com/hay-kot/httpkit/errchain"
	"github.com/signintech/gopdf"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"image/png"
	"io"
	"net/http"

	_ "embed"
)

//go:embed assets/QRIcon.png
var qrcodeLogo []byte

type query struct {
	// 4,296 characters is the maximum length of a QR code
	Data string `schema:"data" validate:"required,max=4296"`
}

type pageQuery struct {
	Nested    *bool  `schema:"nested" validate:"required"`
	PrintType string `schema:"printType" validate:"required,oneof=items locations both"`
	BaseAddr  string `schema:"baseAddr" validate:"required"`
}

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error {
	return nil
}

// HandleGenerateQRCode godoc
//
//	@Summary  Create QR Code
//	@Tags     Items
//	@Produce  json
//	@Param    data      query    string   false "data to be encoded into qrcode"
//	@Success 200 {string} string "image/jpeg"
//	@Router   /v1/qrcode [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleGenerateQRCode() errchain.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) error {
		q, err := adapters.DecodeQuery[query](r)
		if err != nil {
			return err
		}

		image, err := png.Decode(bytes.NewReader(qrcodeLogo))
		if err != nil {
			panic(err)
		}

		qrc, err := qrcode.New(q.Data)
		if err != nil {
			return err
		}

		toWriteCloser := struct {
			io.Writer
			io.Closer
		}{
			Writer: w,
			Closer: io.NopCloser(nil),
		}

		qrwriter := standard.NewWithWriter(toWriteCloser, standard.WithLogoImage(image))

		// Return the QR code as a jpeg image
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", "attachment; filename=qrcode.jpg")
		return qrc.Save(qrwriter)
	}
}

// HandleGenerateQRCodeForLocations godoc
//
// @Summary Create PDF of QR codes of a location
// @Tags Items
// @Produce pdf
// @Param location_id path string true "UUID of the location"
// @Param page query PageQuery true "query info for how the QR codes are generated"
// @Success 200 {string} string "application/pdf"
// @Router /v1/qrcode/{location_id} [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleGenerateQRCodeForLocations() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {

		data, err := adapters.DecodeQuery[pageQuery](r) //decode the query with the pageQuery struct
		if err != nil {
			return err
		}

		routeUUID, err := ctrl.routeUUID(r, "location_id") //from /qrcode/{location_id}
		if err != nil {
			return err
		}

		err = validator.New().Var(routeUUID.String(), "required,uuid") //validate the UUID, so .Get below won't throw errors
		if err != nil {
			return err
		}

		fmt.Println(routeUUID)

		auth := services.NewContext(r.Context())
		locations, err := ctrl.repo.Locations.Get(auth, routeUUID)

		if err != nil {
			return err
		}

		var URLs []string

		if *data.Nested {

		}

		if data.PrintType == "items" {
			for _, thing := range locations.Items {
				encodeStr := data.BaseAddr + "/item/" + thing.ID.String() //concat the url and then UUID
				URLs = append(URLs, encodeStr)
			}
		} else if data.PrintType == "locations" {
			for _, thing := range locations.Children {
				encodeStr := data.BaseAddr + "/location/" + thing.ID.String() //concat the url and then UUID
				URLs = append(URLs, encodeStr)
			}
		} else if data.PrintType == "both" {
			for _, thing := range locations.Items {
				encodeStr := data.BaseAddr + "/item/" + thing.ID.String() //concat the url and then UUID
				URLs = append(URLs, encodeStr)
			}
			for _, thing := range locations.Children {
				encodeStr := data.BaseAddr + "/location/" + thing.ID.String() //concat the url and then UUID
				URLs = append(URLs, encodeStr)
			}
		}

		pdf := gopdf.GoPdf{}
		pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
		pdf.AddPage()
		rowCounter := 0

		const qrPerRow = 4
		const qrPerPage = 24

		for i, thing := range URLs {
			if i%qrPerRow == 0 && i != 0 {
				rowCounter++
				if i%qrPerPage == 0 {
					pdf.AddPage()
				}
			}

			//turn each thing into a image buffer that gopdf accepts
			qrc, err := qrcode.New(thing)
			if err != nil {
				return err
			}

			logo, err := png.Decode(bytes.NewReader(qrcodeLogo))
			if err != nil {
				return err
			}

			img := bytes.NewBuffer(nil)
			wr := standard.NewWithWriter(nopCloser{Writer: img}, standard.WithLogoImage(logo), standard.WithBorderWidth(5), standard.WithQRWidth(7))

			err = qrc.Save(wr)
			if err != nil {
				return err
			}

			//add the image buffer to the pdf
			imgBytes, err := gopdf.ImageHolderByBytes(img.Bytes())
			if err != nil {
				return err
			}

			x := (i % 4) * 150
			y := rowCounter * 150
			err = pdf.ImageByHolder(imgBytes, float64(x), float64(y), nil)
			if err != nil {
				return err
			}

		}

		// Return the concatenated QR code images as a response
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "inline; filename=qrCodes.pdf")

		fmt.Printf("%v", URLs)
		//_, err = w.Write([]byte(fmt.Sprintf("%v", URLs)))
		_, err = pdf.WriteTo(w)
		if err != nil {
			return err
		}
		return err
	}
}
