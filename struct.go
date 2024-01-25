package main

import (
	"encoding/xml"
)

// Components was generated 2024-01-24 19:45:11 by sweet on sweetd.
type Components struct {
	XMLName   xml.Name `xml:"components"`
	Text      string   `xml:",chardata"`
	Version   string   `xml:"version,attr"`
	Origin    string   `xml:"origin,attr"`
	Component []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		ID   string `xml:"id"` // app.authpass.AuthPass, ap...
		Name []struct {
			Text string `xml:",chardata"` // AuthPass, BlueBubbles, Ca...
			Lang string `xml:"lang,attr"`
		} `xml:"name"`
		Summary []struct {
			Text string `xml:",chardata"` // Password Manager: Keep yo...
			Lang string `xml:"lang,attr"`
		} `xml:"summary"`
		Description []struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
			P    []struct {
				Text string   `xml:",chardata"` // Easily and securely keep ...
				Br   []string `xml:"br"`
			} `xml:"p"`
			Ul []struct {
				Text string   `xml:",chardata"`
				Li   []string `xml:"li"` // All your passwords in one...
			} `xml:"ul"`
			Ol struct {
				Text string   `xml:",chardata"`
				Li   []string `xml:"li"` // Open Flatseal, Select ATL...
			} `xml:"ol"`
		} `xml:"description"`
		Icon []struct {
			Text   string `xml:",chardata"` // app.authpass.AuthPass.png...
			Type   string `xml:"type,attr"`
			Height string `xml:"height,attr"`
			Width  string `xml:"width,attr"`
		} `xml:"icon"`
		Categories struct {
			Text     string   `xml:",chardata"`
			Category []string `xml:"category"` // Security, Utility, Chat, ...
		} `xml:"categories"`
		Kudos struct {
			Text string   `xml:",chardata"`
			Kudo []string `xml:"kudo"` // HiDpiIcon, HiDpiIcon, HiD...
		} `xml:"kudos"`
		ProjectLicense string `xml:"project_license"` // GPL-3.0-or-later, Apache-...
		URL            []struct {
			Text string `xml:",chardata"` // https://github.com/authpa...
			Type string `xml:"type,attr"`
		} `xml:"url"`
		Screenshots struct {
			Text       string `xml:",chardata"`
			Screenshot []struct {
				Text  string `xml:",chardata"`
				Type  string `xml:"type,attr"`
				Image []struct {
					Text   string `xml:",chardata"` // https://data.authpass.app...
					Type   string `xml:"type,attr"`
					Height string `xml:"height,attr"`
					Width  string `xml:"width,attr"`
				} `xml:"image"`
				Caption string `xml:"caption"` // Main window, Main window
			} `xml:"screenshot"`
		} `xml:"screenshots"`
		ContentRating []struct {
			Text             string `xml:",chardata"`
			Type             string `xml:"type,attr"`
			ContentAttribute []struct {
				Text string `xml:",chardata"` // moderate, moderate, moder...
				ID   string `xml:"id,attr"`
			} `xml:"content_attribute"`
		} `xml:"content_rating"`
		Releases struct {
			Text    string `xml:",chardata"`
			Release []struct {
				Text        string `xml:",chardata"`
				Timestamp   string `xml:"timestamp,attr"`
				Version     string `xml:"version,attr"`
				Type        string `xml:"type,attr"`
				Urgency     string `xml:"urgency,attr"`
				Description []struct {
					Text string   `xml:",chardata"`
					Lang string   `xml:"lang,attr"`
					P    []string `xml:"p"` // This release fixes issues...
					Ul   []struct {
						Text string   `xml:",chardata"`
						Li   []string `xml:"li"` // Complete rewriting of lar...
					} `xml:"ul"`
					Ol []struct {
						Text string   `xml:",chardata"`
						Li   []string `xml:"li"` // Better transition when op...
					} `xml:"ol"`
				} `xml:"description"`
				URL struct {
					Text string `xml:",chardata"` // https://gitlab.gnome.org/...
					Type string `xml:"type,attr"`
				} `xml:"url"`
				Location string `xml:"location"` // https://github.com/lxgr-l...
				Size     struct {
					Text string `xml:",chardata"` // 7858728, 9876592, 9876592...
					Type string `xml:"type,attr"`
				} `xml:"size"`
				Checksum struct {
					Text string `xml:",chardata"` // 1f52bc1229b2bcbab58c8be72...
					Type string `xml:"type,attr"`
				} `xml:"checksum"`
			} `xml:"release"`
		} `xml:"releases"`
		Launchable []struct {
			Text string `xml:",chardata"` // app.authpass.AuthPass.des...
			Type string `xml:"type,attr"`
		} `xml:"launchable"`
		Bundle struct {
			Text    string `xml:",chardata"` // app/app.authpass.AuthPass...
			Type    string `xml:"type,attr"`
			Runtime string `xml:"runtime,attr"`
			Sdk     string `xml:"sdk,attr"`
		} `xml:"bundle"`
		DeveloperName []struct {
			Text string `xml:",chardata"` // BlueBubbles, Jan Martin R...
			Lang string `xml:"lang,attr"`
		} `xml:"developer_name"`
		Metadata struct {
			Text  string `xml:",chardata"`
			Value []struct {
				Text string `xml:",chardata"` // true, 1683481241, website...
				Key  string `xml:"key,attr"`
			} `xml:"value"`
		} `xml:"metadata"`
		Translation []struct {
			Text string `xml:",chardata"` // cantara, blurble, dialect...
			Type string `xml:"type,attr"`
		} `xml:"translation"`
		ProjectGroup string `xml:"project_group"` // GNOME, GNOME, GNOME, GNOM...
		Languages    struct {
			Text string `xml:",chardata"`
			Lang []struct {
				Text       string `xml:",chardata"` // cs, da, de, es, eu, fi, f...
				Percentage string `xml:"percentage,attr"`
			} `xml:"lang"`
		} `xml:"languages"`
		Keywords struct {
			Text    string `xml:",chardata"`
			Keyword []struct {
				Text string `xml:",chardata"` // wallpaper, Hintergrund, a...
				Lang string `xml:"lang,attr"`
			} `xml:"keyword"`
		} `xml:"keywords"`
		Provides struct {
			Text string `xml:",chardata"`
			Dbus []struct {
				Text string `xml:",chardata"` // app.drey.Dialect.SearchPr...
				Type string `xml:"type,attr"`
			} `xml:"dbus"`
			Binary   []string `xml:"binary"`   // gmult, warp, cb, rednoteb...
			ID       []string `xml:"id"`       // net.launchpad.gmult, net....
			Modalias []string `xml:"modalias"` // usb:v03F0p0053d*, usb:v03...
			Python3  []string `xml:"python3"`  // libretile-desktop, pysol_...
			Python2  string   `xml:"python2"`  // DisplayCAL
			Library  string   `xml:"library"`  // libnuspell.so.4, libvlc.s...
		} `xml:"provides"`
		Mimetypes struct {
			Text     string   `xml:",chardata"`
			Mimetype []string `xml:"mimetype"` // application/ogg, applicat...
		} `xml:"mimetypes"`
		Custom struct {
			Text  string `xml:",chardata"`
			Value []struct {
				Text string `xml:",chardata"` // true, 1675779966, website...
				Key  string `xml:"key,attr"`
			} `xml:"value"`
		} `xml:"custom"`
		Agreement struct {
			Text             string `xml:",chardata"`
			Type             string `xml:"type,attr"`
			VersionID        string `xml:"version_id,attr"`
			AgreementSection []struct {
				Text        string `xml:",chardata"`
				Description struct {
					Text string   `xml:",chardata"`
					P    []string `xml:"p"` // Welcome to the Lunar Clie...
					Ol   struct {
						Text string   `xml:",chardata"`
						Li   []string `xml:"li"` // Use the Services commerci...
					} `xml:"ol"`
				} `xml:"description"`
			} `xml:"agreement_section"`
		} `xml:"agreement"`
		Requires struct {
			Text   string `xml:",chardata"`
			Memory string `xml:"memory"` // 512, 1024, 8192, 8192
			ID     string `xml:"id"`     // com.discordapp.Discord, n...
			Kernel struct {
				Text    string `xml:",chardata"` // Linux
				Compare string `xml:"compare,attr"`
				Version string `xml:"version,attr"`
			} `xml:"kernel"`
		} `xml:"requires"`
		Vetos struct {
			Text string `xml:",chardata"`
			Veto string `xml:"veto"` // NoDisplay=true, NoDisplay...
		} `xml:"vetos"`
		Suggests []struct {
			Text string   `xml:",chardata"`
			ID   []string `xml:"id"` // info.gnuplot.Gnuplot, com...
		} `xml:"suggests"`
		Pkgname              string `xml:"pkgname"`                // gpredict, scribus, opensh...
		CompulsoryForDesktop string `xml:"compulsory_for_desktop"` // Endless, endless
		Extends              string `xml:"extends"`                // org.gnome.Nautilus.deskto...
	} `xml:"component"`
}
