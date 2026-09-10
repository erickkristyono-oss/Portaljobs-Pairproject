package notification

import (
	"context"
	"fmt"

	"application-service/internal/client"
)

type whatsappNotification struct {
	fonnteClient *client.FonnteClient
}

func NewWhatsAppNotification(
	fonnteClient *client.FonnteClient,
) NotificationService {
	return &whatsappNotification{
		fonnteClient: fonnteClient,
	}
}

func (w *whatsappNotification) SendApplicationApplied(
	ctx context.Context,
	phone string,
	jobTitle string,
) error {

	message := fmt.Sprintf(
		`Halo!

		Lamaran kamu untuk posisi *%s* berhasil dikirim.

		Silakan menunggu proses selanjutnya dari perusahaan.

		Terima kasih.`,
		jobTitle,
	)

	return w.fonnteClient.SendMessage(
		ctx,
		phone,
		message,
	)
}

func (w *whatsappNotification) SendInterviewInvitation(
	ctx context.Context,
	phone string,
	jobTitle string,
	companyName string,
) error {

	message := fmt.Sprintf(
		`Halo!

		Selamat! 🎉

		Kamu mendapatkan undangan interview untuk posisi:

		*%s*

		di perusahaan:

		*%s*

		Silakan cek aplikasi kamu untuk informasi interview selanjutnya.

		Semoga sukses!`,
		jobTitle,
		companyName,
	)

	return w.fonnteClient.SendMessage(
		ctx,
		phone,
		message,
	)
}

func (w *whatsappNotification) SendInterviewSelectedToCompany(
	ctx context.Context,
	phone string,
	candidateName string,
	jobTitle string,
) error {

	message := fmt.Sprintf(
		`Halo!

		Kandidat:

		*%s*

		telah masuk ke tahap interview untuk posisi:

		*%s*

		Silakan cek daftar aplikasi untuk proses selanjutnya.`,
		candidateName,
		jobTitle,
	)

	return w.fonnteClient.SendMessage(
		ctx,
		phone,
		message,
	)
}

func (w *whatsappNotification) SendApplicationAccepted(
	ctx context.Context,
	phone string,
	jobTitle string,
	companyName string,
) error {

	message := fmt.Sprintf(
		`Halo!

		Selamat! 🎉

		Lamaran kamu untuk posisi:

		*%s*

		di:

		*%s*

		telah *DITERIMA*.

		Selamat bergabung dan semoga sukses!`,
		jobTitle,
		companyName,
	)

	return w.fonnteClient.SendMessage(
		ctx,
		phone,
		message,
	)
}

func (w *whatsappNotification) SendApplicationAcceptedToCompany(
	ctx context.Context,
	phone string,
	candidateName string,
	jobTitle string,
) error {

	message := fmt.Sprintf(
		`Halo!

		Kandidat:

		*%s*

		telah menerima tawaran untuk posisi:

		*%s*

		Terima kasih.`,
		candidateName,
		jobTitle,
	)

	return w.fonnteClient.SendMessage(
		ctx,
		phone,
		message,
	)
}

func (w *whatsappNotification) SendApplicationRejected(
	ctx context.Context,
	phone string,
	jobTitle string,
) error {

	message := fmt.Sprintf(
		`Halo!

		Terima kasih telah melamar posisi:

		*%s*

		Setelah proses seleksi, lamaran kamu belum dapat dilanjutkan ke tahap berikutnya.

		Tetap semangat dan jangan berhenti mencari kesempatan baru! 💪`,
		jobTitle,
	)

	return w.fonnteClient.SendMessage(
		ctx,
		phone,
		message,
	)
}
