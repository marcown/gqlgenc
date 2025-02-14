// Code generated by github.com/marcown/gqlgenc, DO NOT EDIT.

package client

import (
	"context"
	"encoding/json"
	"example/somelib"

	"github.com/marcown/gqlgenc/client"
	"github.com/marcown/gqlgenc/client/transport"
)

type Client struct {
	Client *client.Client
}

// OPERATION: AsMap
type AsMap struct {
	AsMap string "json:\"asMap\""
}

// INPUT_OBJECT: AsMapInput
type AsMapInput map[string]interface{}

func NewAsMapInput(reqStr string, reqEp Episode) AsMapInput {
	return map[string]interface{}{
		"reqStr": reqStr,
		"reqEp":  reqEp,
	}
}
func (t AsMapInput) WithOptStr(v *string) AsMapInput {
	t["optStr"] = v
	return t
}
func (t AsMapInput) WithOptEp(v *Episode) AsMapInput {
	t["optEp"] = v
	return t
}

// OPERATION: CreatePost
type CreatePost struct {
	Post CreatePost_Post "json:\"post\""
}

// OPERATION: CreatePost.post
type CreatePost_Post struct {
	ID   string "json:\"id\""
	Text string "json:\"text\""
}

// OPERATION: Cyclic1
type Cyclic1 struct {
	Cyclic *Cyclic1_Cyclic "json:\"cyclic\""
}

// OPERATION: Cyclic1.cyclic
type Cyclic1_Cyclic struct {
	Child *Cyclic1_Cyclic_Child "json:\"child\""
}

// OPERATION: Cyclic1.cyclic.child
type Cyclic1_Cyclic_Child struct {
	Child *Cyclic1_Cyclic_Child_Child "json:\"child\""
}

// OPERATION: Cyclic1.cyclic.child.child
type Cyclic1_Cyclic_Child_Child struct {
	Child *Cyclic1_Cyclic_Child_Child_Child "json:\"child\""
}

// OPERATION: Cyclic1.cyclic.child.child.child
type Cyclic1_Cyclic_Child_Child_Child struct {
	ID string "json:\"id\""
}

// OBJECT: Cyclic2_1
type Cyclic2_1 struct {
	ID    string     "json:\"id\""
	Child *Cyclic2_2 "json:\"child\""
}

// OBJECT: Cyclic2_2
type Cyclic2_2 struct {
	ID    string     "json:\"id\""
	Child *Cyclic2_1 "json:\"child\""
}

// ENUM: Episode
type Episode string

const (
	EpisodeNewhope Episode = "NEWHOPE"
	EpisodeEmpire  Episode = "EMPIRE"
	EpisodeJedi    Episode = "JEDI"
)

// ENUM: FooType_hash1
type FooTypeHash1 string

const (
	FooTypeHash1Hash1 FooTypeHash1 = "hash_1"
	FooTypeHash1Hash2 FooTypeHash1 = "hash_2"
)

// OPERATION: GetBooks
type GetBooks struct {
	Books []GetBooks_Books "json:\"books\""
}

// OPERATION: GetBooks.books
type GetBooks_Books struct {
	Typename     string                       "json:\"__typename\""
	Title        string                       "json:\"title\""
	Textbook     *GetBooks_Books_Textbook     "json:\"-\""
	ColoringBook *GetBooks_Books_ColoringBook "json:\"-\""
}

func (t *GetBooks_Books) UnmarshalJSON(data []byte) error {
	type ΞAlias GetBooks_Books
	var r ΞAlias

	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	*t = GetBooks_Books(r)

	switch r.Typename {
	case "ColoringBook":
		var a GetBooks_Books_ColoringBook
		err = json.Unmarshal(data, &a)
		if err != nil {
			return err
		}

		t.ColoringBook = &a
	case "Textbook":
		var a GetBooks_Books_Textbook
		err = json.Unmarshal(data, &a)
		if err != nil {
			return err
		}

		t.Textbook = &a
	}

	return nil
}

// OPERATION: GetBooks.books.ColoringBook
type GetBooks_Books_ColoringBook struct {
	Colors []string "json:\"colors\""
}

// OPERATION: GetBooks.books.Textbook
type GetBooks_Books_Textbook struct {
	Courses []string "json:\"courses\""
}

// OPERATION: GetEpisodes
type GetEpisodes struct {
	Episodes []Episode "json:\"episodes\""
}

// OPERATION: GetMedias
type GetMedias struct {
	Medias []GetMedias_Medias "json:\"medias\""
}

// OPERATION: GetMedias.medias
type GetMedias_Medias struct {
	Typename string                  "json:\"__typename\""
	Image    *GetMedias_Medias_Image "json:\"-\""
	Video    *GetMedias_Medias_Video "json:\"-\""
}

func (t *GetMedias_Medias) UnmarshalJSON(data []byte) error {
	type ΞAlias GetMedias_Medias
	var r ΞAlias

	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	*t = GetMedias_Medias(r)

	switch r.Typename {
	case "Image":
		var a GetMedias_Medias_Image
		err = json.Unmarshal(data, &a)
		if err != nil {
			return err
		}

		t.Image = &a
	case "Video":
		var a GetMedias_Medias_Video
		err = json.Unmarshal(data, &a)
		if err != nil {
			return err
		}

		t.Video = &a
	}

	return nil
}

// OPERATION: GetMedias.medias.Image
type GetMedias_Medias_Image struct {
	Size int64 "json:\"size\""
}

// OPERATION: GetMedias.medias.Video
type GetMedias_Medias_Video struct {
	Duration int64 "json:\"duration\""
}

// OPERATION: GetRoom
type GetRoom struct {
	Room *GetRoom_Room "json:\"room\""
}

// OPERATION: GetRoom.room
type GetRoom_Room struct {
	Name string        "json:\"name\""
	Hash *FooTypeHash1 "json:\"hash\""
}

// OPERATION: GetRoomFragment
type GetRoomFragment struct {
	Room *RoomFragment "json:\"room\""
}

// OPERATION: GetRoomNonNull
type GetRoomNonNull struct {
	RoomNonNull GetRoomNonNull_RoomNonNull "json:\"roomNonNull\""
}

// OPERATION: GetRoomNonNull.roomNonNull
type GetRoomNonNull_RoomNonNull struct {
	Name string "json:\"name\""
}

// OPERATION: Issue8
type Issue8 struct {
	Issue8 *Issue8_Issue8 "json:\"issue8\""
}

// OPERATION: Issue8.issue8
type Issue8_Issue8 struct {
	Foo1 Issue8_Issue8_Foo1  "json:\"foo1\""
	Foo2 *Issue8_Issue8_Foo2 "json:\"foo2\""
}

// OPERATION: Issue8.issue8.foo1
type Issue8_Issue8_Foo1 struct {
	A Issue8_Issue8_Foo1_A "json:\"a\""
}

// OPERATION: Issue8.issue8.foo1.a
type Issue8_Issue8_Foo1_A struct {
	Aa string "json:\"Aa\""
}

// OPERATION: Issue8.issue8.foo2
type Issue8_Issue8_Foo2 struct {
	A Issue8_Issue8_Foo2_A "json:\"a\""
}

// OPERATION: Issue8.issue8.foo2.a
type Issue8_Issue8_Foo2_A struct {
	Aa string "json:\"Aa\""
}

// OPERATION: OptValue1
type OptValue1 struct {
	OptValue1 *bool "json:\"optValue1\""
}

// OPERATION: OptValue2
type OptValue2 struct {
	OptValue2 *bool "json:\"optValue2\""
}

// INPUT_OBJECT: OptionalValue1
type OptionalValue1 struct {
	Value *Value1 "json:\"value\""
}

// INPUT_OBJECT: OptionalValue2
type OptionalValue2 struct {
	Value *Value2 "json:\"value\""
}

// INPUT_OBJECT: PostCreateInput
type PostCreateInput struct {
	Text string "json:\"text\""
}

// OBJECT: RoomFragment
type RoomFragment struct {
	Name string "json:\"name\""
}

// OBJECT: SomeExtraType
type SomeExtraType struct {
	Child SomeExtraTypeChild "json:\"child\""
}

// OBJECT: SomeExtraTypeChild
type SomeExtraTypeChild struct {
	Child SomeExtraTypeChildChild "json:\"child\""
}

// OBJECT: SomeExtraTypeChildChild
type SomeExtraTypeChildChild struct {
	ID string "json:\"id\""
}

// OPERATION: SubscribeMessageAdded
type SubscribeMessageAdded struct {
	MessageAdded SubscribeMessageAdded_MessageAdded "json:\"messageAdded\""
}

// OPERATION: SubscribeMessageAdded.messageAdded
type SubscribeMessageAdded_MessageAdded struct {
	ID string "json:\"id\""
}

// OPERATION: UploadFile
type UploadFile struct {
	UploadFile UploadFile_UploadFile "json:\"uploadFile\""
}

// OPERATION: UploadFile.uploadFile
type UploadFile_UploadFile struct {
	Size int64 "json:\"size\""
}

// OPERATION: UploadFiles
type UploadFiles struct {
	UploadFiles []UploadFiles_UploadFiles "json:\"uploadFiles\""
}

// OPERATION: UploadFiles.uploadFiles
type UploadFiles_UploadFiles struct {
	Size int64 "json:\"size\""
}

// OPERATION: UploadFilesMap
type UploadFilesMap struct {
	UploadFilesMap UploadFilesMap_UploadFilesMap "json:\"uploadFilesMap\""
}

// OPERATION: UploadFilesMap.uploadFilesMap
type UploadFilesMap_UploadFilesMap struct {
	Somefile UploadFilesMap_UploadFilesMap_Somefile "json:\"somefile\""
}

// OPERATION: UploadFilesMap.uploadFilesMap.somefile
type UploadFilesMap_UploadFilesMap_Somefile struct {
	Size int64 "json:\"size\""
}

// INPUT_OBJECT: UploadFilesMapInput
type UploadFilesMapInput struct {
	Somefile transport.Upload "json:\"somefile\""
}

// Pointer helpers
func AsMapInputPtr(v AsMapInput) *AsMapInput {
	return &v
}
func EpisodePtr(v Episode) *Episode {
	return &v
}
func OptionalValue2Ptr(v OptionalValue2) *OptionalValue2 {
	return &v
}
func Value1Ptr(v Value1) *Value1 {
	return &v
}
func Value2Ptr(v Value2) *Value2 {
	return &v
}
func StringPtr(v string) *string {
	return &v
}

const GetRoomDocument = `query GetRoom ($name: String!) {
	room(name: $name) {
		name
		hash
	}
}
`

func (Ξc *Client) GetRoom(ctх context.Context, name string) (*GetRoom, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"name": name,
	}

	{
		var data GetRoom
		res, err := Ξc.Client.Query(ctх, "GetRoom", GetRoomDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetRoomNonNullDocument = `query GetRoomNonNull ($name: String!) {
	roomNonNull(name: $name) {
		name
	}
}
`

func (Ξc *Client) GetRoomNonNull(ctх context.Context, name string) (*GetRoomNonNull, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"name": name,
	}

	{
		var data GetRoomNonNull
		res, err := Ξc.Client.Query(ctх, "GetRoomNonNull", GetRoomNonNullDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetRoomFragmentDocument = `query GetRoomFragment ($name: String!) {
	room(name: $name) {
		... RoomFragment
	}
}
fragment RoomFragment on Chatroom {
	name
}
`

func (Ξc *Client) GetRoomFragment(ctх context.Context, name string) (*GetRoomFragment, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"name": name,
	}

	{
		var data GetRoomFragment
		res, err := Ξc.Client.Query(ctх, "GetRoomFragment", GetRoomFragmentDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetRoomCustomDocument = `query GetRoomCustom ($name: String!) {
	room(name: $name) {
		name
	}
}
`

func (Ξc *Client) GetRoomCustom(ctх context.Context, name string) (*somelib.CustomRoom, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"name": name,
	}

	{
		var data somelib.CustomRoom
		res, err := Ξc.Client.Query(ctх, "GetRoomCustom", GetRoomCustomDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetMediasDocument = `query GetMedias {
	medias {
		__typename
		... on Image {
			size
		}
		... on Video {
			duration
		}
	}
}
`

func (Ξc *Client) GetMedias(ctх context.Context) (*GetMedias, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{}

	{
		var data GetMedias
		res, err := Ξc.Client.Query(ctх, "GetMedias", GetMediasDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetBooksDocument = `query GetBooks {
	books {
		__typename
		title
		... on Textbook {
			courses
		}
		... on ColoringBook {
			colors
		}
	}
}
`

func (Ξc *Client) GetBooks(ctх context.Context) (*GetBooks, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{}

	{
		var data GetBooks
		res, err := Ξc.Client.Query(ctх, "GetBooks", GetBooksDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const SubscribeMessageAddedDocument = `subscription SubscribeMessageAdded {
	messageAdded(roomName: "test") {
		id
	}
}
`

type MessageSubscribeMessageAdded struct {
	Data       *SubscribeMessageAdded
	Error      error
	Extensions transport.RawExtensions
}

func (Ξc *Client) SubscribeMessageAdded(ctх context.Context) (<-chan MessageSubscribeMessageAdded, func()) {
	Ξvars := map[string]interface{}{}

	{
		res := Ξc.Client.Subscription(ctх, "SubscribeMessageAdded", SubscribeMessageAddedDocument, Ξvars)

		ch := make(chan MessageSubscribeMessageAdded)

		go func() {
			for res.Next() {
				opres := res.Get()

				var msg MessageSubscribeMessageAdded
				if len(opres.Errors) > 0 {
					msg.Error = opres.Errors
				}

				err := opres.UnmarshalData(&msg.Data)
				if err != nil && msg.Error == nil {
					msg.Error = err
				}

				msg.Extensions = opres.Extensions

				ch <- msg
			}

			if err := res.Err(); err != nil {
				ch <- MessageSubscribeMessageAdded{
					Error: err,
				}
			}
			close(ch)
		}()

		return ch, res.Close
	}
}

const CreatePostDocument = `mutation CreatePost ($input: PostCreateInput!) {
	post(input: $input) {
		id
		text
	}
}
`

func (Ξc *Client) CreatePost(ctх context.Context, input PostCreateInput) (*CreatePost, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"input": input,
	}

	{
		var data CreatePost
		res, err := Ξc.Client.Mutation(ctх, "CreatePost", CreatePostDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const UploadFileDocument = `mutation UploadFile ($file: Upload!) {
	uploadFile(file: $file) {
		size
	}
}
`

func (Ξc *Client) UploadFile(ctх context.Context, file transport.Upload) (*UploadFile, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"file": file,
	}

	{
		var data UploadFile
		res, err := Ξc.Client.Mutation(ctх, "UploadFile", UploadFileDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const UploadFilesDocument = `mutation UploadFiles ($files: [Upload!]!) {
	uploadFiles(files: $files) {
		size
	}
}
`

func (Ξc *Client) UploadFiles(ctх context.Context, files []*transport.Upload) (*UploadFiles, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"files": files,
	}

	{
		var data UploadFiles
		res, err := Ξc.Client.Mutation(ctх, "UploadFiles", UploadFilesDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const UploadFilesMapDocument = `mutation UploadFilesMap ($files: UploadFilesMapInput!) {
	uploadFilesMap(files: $files) {
		somefile {
			size
		}
	}
}
`

func (Ξc *Client) UploadFilesMap(ctх context.Context, files UploadFilesMapInput) (*UploadFilesMap, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"files": files,
	}

	{
		var data UploadFilesMap
		res, err := Ξc.Client.Mutation(ctх, "UploadFilesMap", UploadFilesMapDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const Issue8Document = `query Issue8 {
	issue8 {
		foo1 {
			a {
				Aa
			}
		}
		foo2 {
			a {
				Aa
			}
		}
	}
}
`

func (Ξc *Client) Issue8(ctх context.Context) (*Issue8, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{}

	{
		var data Issue8
		res, err := Ξc.Client.Query(ctх, "Issue8", Issue8Document, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const GetEpisodesDocument = `query GetEpisodes {
	episodes
}
`

func (Ξc *Client) GetEpisodes(ctх context.Context) (*GetEpisodes, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{}

	{
		var data GetEpisodes
		res, err := Ξc.Client.Query(ctх, "GetEpisodes", GetEpisodesDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const Cyclic1Document = `query Cyclic1 {
	cyclic {
		child {
			child {
				child {
					id
				}
			}
		}
	}
}
`

func (Ξc *Client) Cyclic1(ctх context.Context) (*Cyclic1, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{}

	{
		var data Cyclic1
		res, err := Ξc.Client.Query(ctх, "Cyclic1", Cyclic1Document, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const AsMapDocument = `query AsMap ($req: AsMapInput!, $opt: AsMapInput) {
	asMap(req: $req, opt: $opt)
}
`

func (Ξc *Client) AsMap(ctх context.Context, req AsMapInput, opt *AsMapInput) (*AsMap, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"req": req,
		"opt": opt,
	}

	{
		var data AsMap
		res, err := Ξc.Client.Query(ctх, "AsMap", AsMapDocument, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const OptValue1Document = `query OptValue1 ($v: OptionalValue1!) {
	optValue1(req: $v)
}
`

func (Ξc *Client) OptValue1(ctх context.Context, v OptionalValue1) (*OptValue1, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"v": v,
	}

	{
		var data OptValue1
		res, err := Ξc.Client.Query(ctх, "OptValue1", OptValue1Document, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}

const OptValue2Document = `query OptValue2 ($v: OptionalValue2) {
	optValue2(opt: $v)
}
`

func (Ξc *Client) OptValue2(ctх context.Context, v *OptionalValue2) (*OptValue2, transport.OperationResponse, error) {
	Ξvars := map[string]interface{}{
		"v": v,
	}

	{
		var data OptValue2
		res, err := Ξc.Client.Query(ctх, "OptValue2", OptValue2Document, Ξvars, &data)
		if err != nil {
			return nil, transport.OperationResponse{}, err
		}

		return &data, res, nil
	}
}
